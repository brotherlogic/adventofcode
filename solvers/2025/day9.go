package main

import (
	"context"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func (s *Server) Day9Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	best := int32(0)

	return &pb.SolveResponse{Answer: best}, nil
}
