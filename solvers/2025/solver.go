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

	return ret[0].Interface().(*pb.SolveResponse), ret[1].Interface().(error)
}

func (s *Server) heartbeat(ctx context.Context) error {
	conn, err := grpc.NewClient(AOC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewAdventOfCodeInternalServiceClient(conn)
	_, err = client.Register(ctx, &pb.RegisterRequest{
		Year:     2025,
		Callback: "adventofcode-solver-2025.adventofcode:8080",
	})
	return err
}

func main() {
	server := &Server{}

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
