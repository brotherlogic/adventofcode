package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	ghb_client "github.com/brotherlogic/githubridge/client"
	rstore_client "github.com/brotherlogic/rstore/client"

	pb "github.com/brotherlogic/adventofcode/proto"
	ghbpb "github.com/brotherlogic/githubridge/proto"
	rspb "github.com/brotherlogic/rstore/proto"
)

var (
	solvingDuration = time.Minute * 5
	retries         = 3
)

func solve(year, day, part int32) error {

	log.Printf("Solving %v %v %v", year, day, part)
	for i := 0; i < retries; i++ {
		err := solveInternal(year, day, part)
		if status.Code(err) != codes.NotFound {
			return err
		}
	}

	return status.Errorf(codes.ResourceExhausted, "Unable to solve with retries")
}

func solveInternal(year, day, part int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), solvingDuration)
	defer cancel()

	conn, err := grpc.Dial("adventofcode.adventofcode:8080", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}

	client := pb.NewSolverServiceClient(conn)
	iclient := pb.NewAdventOfCodeInternalServiceClient(conn)
	res, err := client.Solve(ctx, &pb.SolveRequest{
		Year: year,
		Day:  day,
		Part: part,
	})

	if err != nil {
		return fmt.Errorf("bad solve: %w", err)
	}

	sol, err := iclient.GetSolution(ctx, &pb.GetSolutionRequest{
		Year: year,
		Day:  day,
		Part: part,
	})
	if status.Code(err) == codes.OK {
		if sol.GetSolution().GetBigAnswer() == res.GetBigAnswer() {
			return nil
		}

		return status.Errorf(codes.FailedPrecondition, "Solution is not present or incorrect %v vs %v", sol.GetSolution(), res)
	}

	return err
}

func loadExistingIssue(ctx context.Context, rsclient rstore_client.RStoreClient) (int32, error) {
	data, err := rsclient.Read(ctx, &rspb.ReadRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
	if err != nil {
		return -1, nil
	}
	return int32(binary.LittleEndian.Uint32(data.GetValue().GetValue())), nil
}

func raiseIssue(ctx context.Context, ghclient ghb_client.GithubridgeClient, rsclient rstore_client.RStoreClient, year, day, part int, err error) error {
	issue, err := ghclient.AddIssue(ctx, &ghbpb.AddIssueRequest{Title: fmt.Sprintf("Solve %v - %v - %v (%v)", year, day, part, err), Job: "adventofcode"})
	if err != nil {
		return err
	}

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(issue.GetIssueId()))
	_, err = rsclient.Write(ctx, &rspb.WriteRequest{Key: "brotherlogic/adventofcode/finder/cissue", Value: &anypb.Any{Value: b}})
	return err
}

func findIssue(iid int32) error {
	return nil
}

func runYear(ctx context.Context, ghclient ghb_client.GithubridgeClient, rsclient rstore_client.RStoreClient, year, db int) error {
	for day := 1; day <= db; day++ {
		for part := 1; part <= 2; part++ {
			err := solve(int32(year), int32(day), int32(part))
			log.Printf("Solved %v %v %v -> %v", year, day, part, err)
			if status.Code(err) != codes.OK {
				//Raise the issue to solve this problem
				return raiseIssue(ctx, ghclient, rsclient, year, day, part, err)
			}
		}
	}

	return nil
}

func main() {
	log.Print("Running")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ghclient, err := ghb_client.GetClient()
	if err != nil {
		log.Fatalf("unable to get ghb client: %v", err)
	}

	rstore, err := rstore_client.GetClient()
	if err != nil {
		log.Fatalf("unable to get rstore client: %v", err)
	}

	// Check on the existing issue
	iid, err := loadExistingIssue(ctx, rstore)
	if err != nil {
		log.Fatalf("unable to load existing issue: %v", err)
	}

	// We have no solved the current issue
	if iid > 0 {
		log.Printf("Issue is still open: %v", iid)
		return
	}

	// If we're in a set, run this
	if time.Now().Month() == time.December && time.Now().Day() <= 25 {
		err = runYear(ctx, ghclient, rstore, time.Now().Year(), time.Now().Day())
		log.Printf("Result: %v", err)
		return
	}

	// If we're not in a set, work days at a time
	for day := 1; day <= 25; day++ {
		for year := 2015; year < time.Now().Year(); year++ {
			if runYear(ctx, ghclient, rstore, year, day) != nil {
				return
			}
		}
	}

	log.Printf("No more puizzles to solve")
}
