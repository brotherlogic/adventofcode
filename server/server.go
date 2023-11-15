package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	pb "github.com/brotherlogic/adventofcode/proto"
	rspb "github.com/brotherlogic/rstore/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type Server struct {
	years   map[int32]bool
	solvers map[string]bool
	gs      *grpc.Server
}

func (s *Server) Run(pn int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", pn))
	if err != nil {
		return err
	}

	s.gs = grpc.NewServer()
	pb.RegisterAdventOfCodeServiceServer()
}

func (s *Server) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	conn, err := grpc.Dial("rstore.rstore:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := rspb.NewRStoreServiceClient(conn)

	_, err = client.Write(ctx, &rspb.WriteRequest{
		Key:   fmt.Sprintf("adventofcode/data/%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart()),
		Value: &anypb.Any{Value: []byte(req.GetData())},
	})

	return &pb.UploadResponse{}, err
}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	wg := &sync.WaitGroup{}

	var resp *pb.SolveResponse
	for callback, _ := range s.solvers {
		wg.Add(1)
		go func() {
			defer wg.Done()

			conn, err := grpc.Dial(callback, grpc.WithInsecure())
			if err != nil {
				return
			}
			defer conn.Close()

			client := pb.NewSolverServiceClient(conn)
			res, err := client.Solve(ctx, req)
			if err != nil {
				return
			}

			resp = res
		}()
	}

	if resp == nil {
		return resp, status.Errorf(codes.NotFound, "unable to find solver for %v", req)
	}

	return resp, nil
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	for _, year := range req.GetYears() {
		s.years[year] = true
	}
	s.solvers[req.GetCallback()] = true
	return &pb.RegisterResponse{}, nil
}
