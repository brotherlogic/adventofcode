package server

import (
	"context"
	"fmt"

	pb "github.com/brotherlogic/adventofcode/proto"
	rspb "github.com/brotherlogic/rstore/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type Server struct{}

func (s *Server) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	conn, err := grpc.Dial("rstore.rstore:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := rspb.NewRStoreServiceClient(conn)

	_, err = client.Write(ctx, &rspb.WriteRequest{
		Key:   fmt.Sprintf("adventofcode/data/%v-%v-%v", req.GetYear(), req.GetDay(), req.GetPart()),
		Value: &anypb.Any{Value: []byte(req.GetData())},
	})

	return &pb.UploadResponse{}, err
}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Haven't done this yet")
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Haven't done this yet")
}
