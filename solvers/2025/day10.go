package main

import (
	"context"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func (s *Server) Day10Part1(_ context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int32(0)

	return &pb.SolveResponse{Answer: sumv}, nil
}
