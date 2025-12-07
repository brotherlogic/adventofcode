package main

import (
	"context"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func (s *Server) Day6Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int64(0)

	return &pb.SolveResponse{
		BigAnswer: sumv,
	}, nil
}
