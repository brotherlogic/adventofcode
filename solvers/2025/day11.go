package main

import (
	"context"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func (*Server) Day11Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	count := int32(0)
	return &pb.SolveResponse{Answer: count}, nil
}
