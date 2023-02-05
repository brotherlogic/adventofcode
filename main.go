package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/brotherlogic/adventofcode/server"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/adventofcode/proto"
)

var (
	port = flag.Int("port", 8080, "The server port.")
)

func main() {
	s := &server.Server{}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	pb.RegisterAdventServerServiceServer(gs, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
