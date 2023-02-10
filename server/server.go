package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/adventofcode/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	//Backlog - the print queue
	count = promauto.NewCounter(prometheus.CounterOpts{
		Name: "adventofcode_count",
		Help: "The number of overall requests",
	})
)

type Server struct{}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	count.Inc()
	return nil, status.Errorf(codes.Unimplemented, "We actually haven't written this yet")
}
