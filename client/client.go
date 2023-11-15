package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/brotherlogic/adventofcode/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewAdventOfCodeServiceClient(conn)
	res, err := client.Solve(ctx, &pb.SolveRequest{Year: 2017, Day: 10, Part: 2})

	fmt.Printf("%v -> %v\n", res, err)
}
