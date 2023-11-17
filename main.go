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
	port         = flag.Int("port", 8080, "The server port.")
	metricsPort  = flag.Int("metrics_port", 8081, "Metrics port")
	internalPort = flag.Int("internal_port", 8082, "Internal services port")
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

	lisInternal, err := net.Listen("tcp", fmt.Sprintf(":%d", *internalPort))
	if err != nil {
		log.Fatalf("failed to listen on internal port %v: %v", *internalPort, err)
	}
	gsi := grpc.NewServer()
	pb.RegisterAdventOfCodeInternalServiceServer(gsi, s)

	// Setup prometheus export
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
	}()

	go func() {
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
	}()

	log.Printf("Serving aoc hub")
	if err := gsi.Serve(lisInternal); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
