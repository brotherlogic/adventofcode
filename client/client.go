package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/brotherlogic/adventofcode/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewAdventServerServiceClient(conn)
	res, err := client.Solve(context.Background(), &pb.SolveRequest{})

	fmt.Printf("%v -> %v\n", res, err)
}
