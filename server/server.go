package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	pb "github.com/brotherlogic/adventofcode/proto"
	rspb "github.com/brotherlogic/rstore/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type Server struct {
	years   map[int32]bool
	solvers map[string]bool
}

func NewServer() *Server {
	return &Server{
		years:   make(map[int32]bool),
		solvers: make(map[string]bool),
	}
}

var (
	years = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adventofcode_years",
		Help: "The size of the print queue",
	}, []string{"year"})
)

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
	conn, err := grpc.Dial("rstore.rstore:8080", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot dial rstore: %w", err)
	}

	client := rspb.NewRStoreServiceClient(conn)
	resp, err := client.Read(ctx, &rspb.ReadRequest{
		Key: fmt.Sprintf("adventofcode/data/%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart()),
	})
	if err != nil {
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
		return solution, nil
	}

	if len(errors) == 0 {
		return nil, status.Errorf(codes.Unimplemented, "No solvers for %v/%v/%v", req.GetYear(), req.GetDay(), req.GetPart())
	}

	return nil, status.Errorf(codes.Internal, "Many errors: %v", errors)
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.years[req.GetYear()] = true
	s.solvers[req.GetCallback()] = true

	log.Printf("Received and stored: %v", req)

	return &pb.RegisterResponse{}, s.updateMetrics()
}
