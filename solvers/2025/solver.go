package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"reflect"
	"time"

	pb "github.com/brotherlogic/adventofcode/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	AOC_ADDRESS = "adventofcode.adventofcode:8082"
)

var (
	port = flag.Int("port", 8080, "The server port.")
)

type Server struct {
	startup int64
}

func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	day := req.GetDay()
	part := req.GetPart()

	methodName := fmt.Sprintf("Day%vPart%v", day, part)
	log.Printf("Trying to run %v", methodName)
	method := reflect.ValueOf(&Server{}).MethodByName(methodName)
	if !method.IsValid() {
		return nil, fmt.Errorf("cannot find method Day%vPart%v", day, part)
	}

	ret := method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(req),
	})

	if ret[0].IsNil() {
		return nil, ret[1].Interface().(error)
	}

	return ret[0].Interface().(*pb.SolveResponse), nil
}

func (s *Server) heartbeat(ctx context.Context) error {
	conn, err := grpc.NewClient(AOC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewAdventOfCodeInternalServiceClient(conn)
	_, err = client.Register(ctx, &pb.RegisterRequest{
		Year:        2025,
		Callback:    "solver-2025.adventofcode:8080",
		StartupTime: s.startup,
	})
	return err
}

func main() {
	server := &Server{
		startup: time.Now().Unix(),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", *port, err)
	}
	gs := grpc.NewServer()
	pb.RegisterSolverServiceServer(gs, server)

	// Run a heartbeat every minute
	go func() {
		for {
			time.Sleep(time.Minute)

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			err := server.heartbeat(ctx)
			if err != nil {
				log.Printf("Unable to send heartbeat: %v", err)
			}
			cancel()
		}
	}()

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
