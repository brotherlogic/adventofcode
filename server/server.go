package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/brotherlogic/adventofcode/proto"
	pspb "github.com/brotherlogic/pstore/proto"

	pstore_client "github.com/brotherlogic/pstore/client"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Server struct {
	years    map[int32]bool
	solvers  map[string]int32
	psclient pstore_client.PStoreClient
	mapLock  *sync.Mutex
}

func NewServer(psclient pstore_client.PStoreClient) *Server {
	return &Server{
		years:    make(map[int32]bool),
		solvers:  make(map[string]int32),
		psclient: psclient,
		mapLock:  &sync.Mutex{},
	}
}

var (
	years = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adventofcode_years",
		Help: "The size of the print queue",
	}, []string{"year"})

	solveRequest = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adventofcode_solves",
		Help: "The size of the print queue",
	}, []string{"puzzle", "result"})

	solveTimes = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "adventofcode_times",
		Help:    "The size of the print queue",
		Buckets: []float64{1, 10, 100, 1000, 2000, 4000, 8000, 16000, 32000, 64000},
	}, []string{"puzzle", "result"})
)

func (s *Server) GetSolution(ctx context.Context, req *pb.GetSolutionRequest) (*pb.GetSolutionResponse, error) {
	data, err := s.psclient.Read(ctx, &pspb.ReadRequest{Key: "github.com/brotherlogic/adventofcode/solutions"})
	if err != nil {
		return nil, err
	}

	solutions := &pb.Solutions{}
	err = proto.Unmarshal(data.GetValue().GetValue(), solutions)
	if err != nil {
		return nil, err
	}

	solresp := &pb.GetSolutionResponse{Solution: &pb.Solution{}}
	found := false
	for _, solution := range solutions.GetSolutions() {
		if solution.GetDay() == req.GetDay() && solution.GetYear() == req.GetYear() && solution.GetPart() == req.GetPart() {
			if solution.GetAnswer() > 0 {
				solresp.Solution.Answer = solution.GetAnswer()
				found = true
			}
			if solution.GetBigAnswer() > 0 {
				solresp.Solution.BigAnswer = solution.GetBigAnswer()
				found = true
			}
			if solution.GetStringAnswer() != "" {
				solresp.Solution.BigAnswer = solution.GetBigAnswer()
				found = true
			}
		}
	}
	if found {
		return solresp, nil
	}

	return nil, status.Errorf(codes.NotFound, "Unable to locate solution for %v %v %v", req.GetYear(), req.GetDay(), req.GetPart())
}

func (s *Server) AddSolution(ctx context.Context, req *pb.AddSolutionRequest) (*pb.AddSolutionResponse, error) {
	data, err := s.psclient.Read(ctx, &pspb.ReadRequest{Key: "github.com/brotherlogic/adventofcode/solutions"})
	if err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}

	solutions := &pb.Solutions{Solutions: make([]*pb.Solution, 0)}
	if status.Code(err) != codes.NotFound {
		err = proto.Unmarshal(data.GetValue().GetValue(), solutions)
		if err != nil {
			return nil, err
		}
	}

	solutions.Solutions = append(solutions.Solutions, req.GetSolution())
	ndata, err := proto.Marshal(solutions)
	if err != nil {
		return nil, err
	}

	_, err = s.psclient.Write(ctx, &pspb.WriteRequest{
		Key:   "github.com/brotherlogic/adventofcode/solutions",
		Value: &anypb.Any{Value: ndata},
	})
	return &pb.AddSolutionResponse{}, err
}

func (s *Server) updateMetrics() error {
	for year := range s.years {
		years.With(prometheus.Labels{"year": fmt.Sprintf("%v", year)}).Set(1)
	}

	return nil
}

func (s *Server) SetCookie(ctx context.Context, req *pb.SetCookieRequest) (*pb.SetCookieResponse, error) {
	conn, err := grpc.Dial("pstore.pstore:8080", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial error on pstore: %w", err)
	}

	client := pspb.NewPStoreServiceClient(conn)
	_, err = client.Write(ctx, &pspb.WriteRequest{
		Key:   "brotherlogic/adventofcode/finder/cookie",
		Value: &anypb.Any{Value: []byte(req.GetCookie())},
	})
	return &pb.SetCookieResponse{}, err
}

func (s *Server) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	conn, err := grpc.Dial("pstore.pstore:8080", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial error on pstore: %w", err)
	}

	client := pspb.NewPStoreServiceClient(conn)

	_, err = client.Write(ctx, &pspb.WriteRequest{
		Key:   fmt.Sprintf("adventofcode/data/%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart()),
		Value: &anypb.Any{Value: []byte(req.GetData())},
	})

	if err != nil {
		return &pb.UploadResponse{}, fmt.Errorf("bad write: %w", err)
	}

	return &pb.UploadResponse{}, nil
}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	log.Printf("%v with %v", req, s.solvers)
	// Validate existance first
	for _, solver := range s.solvers {
		if solver == req.GetYear() {
			if req.GetDay() == 0 {
				return &pb.SolveResponse{}, nil
			}
			break
		}

		return &pb.SolveResponse{}, status.Errorf(codes.InvalidArgument, "Unable to find solver for %v", req.GetYear())
	}

	puzzle := fmt.Sprintf("%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart())
	conn, err := grpc.Dial("pstore.pstore:8080", grpc.WithInsecure())
	if err != nil {
		solveRequest.With(prometheus.Labels{"puzzle": puzzle, "result": fmt.Sprintf("Dial %v", status.Code(err))}).Inc()
		return nil, fmt.Errorf("cannot dial pstore: %w", err)
	}

	client := pspb.NewPStoreServiceClient(conn)
	resp, err := client.Read(ctx, &pspb.ReadRequest{
		Key: fmt.Sprintf("adventofcode/data/%v-%v-1", req.GetYear(), req.GetDay()),
	})
	if err != nil {
		solveRequest.With(prometheus.Labels{"puzzle": puzzle, "result": fmt.Sprintf("Read: %v", status.Code(err))}).Inc()
		return nil, fmt.Errorf("bad read: %w", err)
	}
	req.Data = string(resp.GetValue().GetValue())

	var errors []error
	var solution *pb.SolveResponse

	wg := &sync.WaitGroup{}
	for callback := range s.solvers {
		conn, err := grpc.Dial(callback, grpc.WithInsecure())
		if err != nil {
			errors = append(errors, err)
			continue
		}
		wg.Add(1)
		defer conn.Close()
		go func(conn *grpc.ClientConn, c string, puzzle string) {
			client := pb.NewSolverServiceClient(conn)
			t1 := time.Now()
			tsol, err := client.Solve(ctx, req)
			log.Printf("Solved %v %v in %v -> %v", puzzle, c, time.Since(t1), err)
			solveTimes.With(prometheus.Labels{
				"puzzle": fmt.Sprintf("%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart()),
				"result": fmt.Sprintf("%v", status.Code(err)),
			}).Observe(float64(time.Since(t1).Milliseconds()))
			wg.Done()
			if err != nil {
				errors = append(errors, fmt.Errorf("%v -> %w", c, err))
				return
			}
			solution = tsol
		}(conn, callback, puzzle)
	}

	wg.Wait()

	if solution != nil {
		solveRequest.With(prometheus.Labels{"puzzle": puzzle, "result": "OK"}).Inc()
		return solution, nil
	}

	if len(errors) == 0 {
		solveRequest.With(prometheus.Labels{"puzzle": puzzle, "result": "Unimplemented"}).Inc()
		return nil, status.Errorf(codes.Unimplemented, "No solvers for %v/%v/%v", req.GetYear(), req.GetDay(), req.GetPart())
	}

	solveRequest.With(prometheus.Labels{"puzzle": puzzle, "result": "Internal"}).Inc()
	return nil, status.Errorf(codes.Internal, "Many errors: %v", errors)
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.mapLock.Lock()
	s.years[req.GetYear()] = true
	s.solvers[req.GetCallback()] = req.GetYear()
	s.mapLock.Unlock()

	return &pb.RegisterResponse{}, s.updateMetrics()
}
