package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/brotherlogic/adventofcode/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 8080, "The server port.")
)

type Server struct {
}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{}, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}

func main() {
	server := &Server{}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", *port, err)
	}
	gs := grpc.NewServer()
	pb.RegisterSolverServiceServer(gs, server)

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
