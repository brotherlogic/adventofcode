package server

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/adventofcode/proto"
	aspb "github.com/brotherlogic/adventserver/proto"
	rspb "github.com/brotherlogic/rstore/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"

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

func (s *Server) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	conn, err := grpc.Dial("rstore.rstore:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := rspb.NewRStoreServiceClient(conn)
	_, err = client.Write(ctx, &rspb.WriteRequest{
		Key:   fmt.Sprintf("adventofcode/data/%v-%v", req.GetYear(), req.GetYear()),
		Value: &anypb.Any{Value: []byte(req.GetData())},
	})

	return &pb.UploadResponse{}, err
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
