package main

import (
	"context"
	"fmt"
	"io/ioutil"
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
	iclient := pb.NewAdventOfCodeInternalServiceClient(conn)

	switch os.Args[2] {
	case "solve":
		res, err := client.Solve(ctx, &pb.SolveRequest{Year: 2015, Day: 1, Part: 1})
		fmt.Printf("%v -> %v\n", res, err)
	case "upload":
		data, err := ioutil.ReadFile(os.Args[3])
		if err != nil {
			log.Fatalf("Unable to run upload: %v", err)
		}
		res, err := iclient.Upload(ctx, &pb.UploadRequest{Year: 2015, Day: 1, Part: 1, Data: string(data)})
		fmt.Printf("%v -> %v\n", res, err)
	default:
		fmt.Printf("Unknown command: %v\n", os.Args[1])
	}
}
