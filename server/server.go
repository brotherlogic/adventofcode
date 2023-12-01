package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "github.com/brotherlogic/adventofcode/proto"
	rspb "github.com/brotherlogic/rstore/proto"

	rstore_client "github.com/brotherlogic/rstore/client"

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
	solvers  map[string]bool
	rsclient rstore_client.RStoreClient
}

func NewServer(rsclient rstore_client.RStoreClient) *Server {
	return &Server{
		years:    make(map[int32]bool),
		solvers:  make(map[string]bool),
		rsclient: rsclient,
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
)

func (s *Server) GetSolution(ctx context.Context, req *pb.GetSolutionRequest) (*pb.GetSolutionResponse, error) {
	data, err := s.rsclient.Read(ctx, &rspb.ReadRequest{Key: "github.com/brotherlogic/adventofcode/solutions"})
	if err != nil {
		return nil, err
	}

	log.Printf("HERE %v %v", data, err)
	solutions := &pb.Solutions{}
	err = proto.Unmarshal(data.GetValue().GetValue(), solutions)
	if err != nil {
		return nil, err
	}

	for _, solution := range solutions.GetSolutions() {
		if solution.GetDay() == req.GetDay() && solution.GetYear() == req.GetYear() && solution.GetPart() == req.GetPart() {
			return &pb.GetSolutionResponse{Solution: solution}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Unable to locate solution for %v %v %v", req.GetYear(), req.GetDay(), req.GetPart())
}

func (s *Server) AddSolution(ctx context.Context, req *pb.AddSolutionRequest) (*pb.AddSolutionResponse, error) {
	data, err := s.rsclient.Read(ctx, &rspb.ReadRequest{Key: "github.com/brotherlogic/adventofcode/solutions"})
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

	_, err = s.rsclient.Write(ctx, &rspb.WriteRequest{
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

func (s *Server) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	conn, err := grpc.Dial("rstore.rstore:8080", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial error on rstore: %w", err)
	}

	client := rspb.NewRStoreServiceClient(conn)

	_, err = client.Write(ctx, &rspb.WriteRequest{
		Key:   fmt.Sprintf("adventofcode/data/%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart()),
		Value: &anypb.Any{Value: []byte(req.GetData())},
	})

	return &pb.UploadResponse{}, fmt.Errorf("bad write: %w", err)
}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	puzzle := fmt.Sprintf("%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart())
	conn, err := grpc.Dial("rstore.rstore:8080", grpc.WithInsecure())
	if err != nil {
		solveRequest.With(prometheus.Labels{"puzzle": puzzle, "result": fmt.Sprintf("Dial %v", status.Code(err))}).Inc()
		return nil, fmt.Errorf("cannot dial rstore: %w", err)
	}

	client := rspb.NewRStoreServiceClient(conn)
	resp, err := client.Read(ctx, &rspb.ReadRequest{
		Key: fmt.Sprintf("adventofcode/data/%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart()),
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
		go func(conn *grpc.ClientConn) {
			client := pb.NewSolverServiceClient(conn)
			tsol, err := client.Solve(ctx, req)
			wg.Done()
			if err != nil {
				errors = append(errors, err)
				return
			}
			solution = tsol
		}(conn)
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
	s.years[req.GetYear()] = true
	s.solvers[req.GetCallback()] = true

	log.Printf("Received and stored: %v", req)

	return &pb.RegisterResponse{}, s.updateMetrics()
}
