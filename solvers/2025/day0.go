package main

import (
	"context"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func (s *Server) Day0Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{Answer: int32(123)}, nil
}
