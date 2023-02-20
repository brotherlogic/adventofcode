package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	count.Inc()

	if req.GetYear() == 2017 && req.GetDay() == 12 && req.GetPart() == 2 {
		conn, err := utils.LFDialServer(ctx, "adventserver")
		if err != nil {
			return nil, err
		}
		client := aspb.NewAdventServerServiceClient(conn)
		val, err := client.Solve(ctx, &aspb.SolveRequest{Year: req.GetYear(), Day: req.GetDay(), Part: req.GetPart()})
		if err != nil {
			return nil, err
		}
		return &pb.SolveResponse{Answer: val.GetAnswer()}, nil
	}

	return nil, status.Errorf(codes.Unimplemented, "we actually haven't written this yet")
}
