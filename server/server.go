package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/adventofcode/proto"
)

type Server struct{}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "we actually haven't written this yet")
}
