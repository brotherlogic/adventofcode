package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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

func solve(ctx context.Context, year, day, part int32, issue *pb.Issue) error {

	log.Printf("Solving %v %v %v", year, day, part)
	for i := 0; i < retries; i++ {
		err := solveInternal(ctx, year, day, part, issue)
		if status.Code(err) != codes.NotFound {
			return err
		}
		log.Printf("Solve fail: %v", err)
	}

	return status.Errorf(codes.ResourceExhausted, "Unable to solve with retries")
}

func solveInternal(sctx context.Context, year, day, part int32, issue *pb.Issue) error {
	ctx, cancel := context.WithTimeout(context.Background(), solvingDuration)
	defer cancel()

	connm, err := grpc.Dial("adventofcode.adventofcode:8080", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}
	conni, err := grpc.Dial("adventofcode.adventofcode:8082", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}

	client := pb.NewAdventOfCodeServiceClient(connm)
	iclient := pb.NewAdventOfCodeInternalServiceClient(conni)
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

	if status.Code(err) == codes.NotFound {
		return addSolutionToIssue(ctx, sol.GetSolution(), issue)
	}

	return err
}

func loadExistingIssue(ctx context.Context, rsclient rstore_client.RStoreClient) (*pb.Issue, error) {
	data, err := rsclient.Read(ctx, &rspb.ReadRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
	if err != nil {
		return nil, err
	}

	issue := &pb.Issue{}
	err = proto.Unmarshal(data.GetValue().GetValue(), issue)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func addSolutionToIssue(ctx context.Context, solution *pb.Solution, issue *pb.Issue) error {
	return nil
}

func raiseIssue(ctx context.Context, ghclient ghb_client.GithubridgeClient, rsclient rstore_client.RStoreClient, year, day, part int, err error) error {
	issue, err := ghclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{Title: fmt.Sprintf("Solve %v - %v - %v (%v)", year, day, part, err), Repo: "adventofcode", User: "brotherlogic"})
	if err != nil {
		return err
	}

	iss := &pb.Issue{
		Id:   issue.GetIssueId(),
		Open: true,
	}
	bytes, err := proto.Marshal(iss)
	if err != nil {
		return err
	}

	_, err = rsclient.Write(ctx, &rspb.WriteRequest{Key: "brotherlogic/adventofcode/finder/cissue", Value: &anypb.Any{Value: bytes}})
	return err
}

func findIssue(iid int32) error {
	return nil
}

func runYear(ctx context.Context, ghclient ghb_client.GithubridgeClient, rsclient rstore_client.RStoreClient, year, db int, issue *pb.Issue) error {
	for day := 1; day <= db; day++ {
		for part := 1; part <= 2; part++ {
			err := solve(ctx, int32(year), int32(day), int32(part), issue)
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
	issue, err := loadExistingIssue(ctx, rstore)
	if err != nil {
		log.Fatalf("unable to load existing issue: %v", err)
	}

	// We have no solved the current issue
	if issue == nil || !issue.GetOpen() {
		log.Printf("Issue is still open: %v", issue)
		return
	}

	// If we're in a set, run this
	if time.Now().Month() == time.December && time.Now().Day() <= 25 {
		err = runYear(ctx, ghclient, rstore, time.Now().Year(), time.Now().Day(), issue)
		log.Printf("Result: %v", err)
		return
	}

	// If we're not in a set, work days at a time
	for day := 1; day <= 25; day++ {
		for year := 2015; year < time.Now().Year(); year++ {
			if runYear(ctx, ghclient, rstore, year, day, issue) != nil {
				return
			}
		}
	}

	log.Printf("No more puizzles to solve")
}
