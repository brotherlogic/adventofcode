package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/brotherlogic/adventofcode/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/adventofcode/proto"
)

var (
	port        = flag.Int("port", 8080, "The server port.")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", *port, err)
	}
	gs := grpc.NewServer()
	pb.RegisterAdventOfCodeServiceServer(gs, s)
	log.Printf("Server listening at %v", lis.Addr())

	// Setup prometheus export
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
	}()

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
