package server

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/adventofcode/proto"
	aspb "github.com/brotherlogic/adventserver/proto"

	"github.com/brotherlogic/goserver/utils"

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

func (s *Server) localSolve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return nil, fmt.Errorf("No local solves yet")
}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	count.Inc()

	resp, err := s.localSolve(ctx, req)
	if err == nil {
		return resp, err
	}

	conn, err := utils.LFDialServer(ctx, "adventserver")
	if err != nil {
		return nil, err
	}
	client := aspb.NewAdventServerServiceClient(conn)
	val, err := client.Solve(ctx, &aspb.SolveRequest{Year: req.GetYear(), Day: req.GetDay(), Part: req.GetPart()})
	if err != nil {
		return nil, err
	}
	return &pb.SolveResponse{Answer: val.GetAnswer(), StringAnswer: val.GetStringAnswer(), BigAnswer: val.GetBigAnswer()}, nil
}
