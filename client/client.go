package main

import (
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/brotherlogic/adventofcode/proto"
	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := utils.ManualContext("adventofcode", time.Minute)
	defer cancel()

	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewAdventServerServiceClient(conn)
	res, err := client.Solve(ctx, &pb.SolveRequest{})

	fmt.Printf("%v -> %v\n", res, err)
}
