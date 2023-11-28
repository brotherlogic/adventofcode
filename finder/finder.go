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

	conn, err := grpc.Dial("adventofocde.adventofcode:8080", grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := pb.NewSolverServiceClient(conn)
	_, err = client.Solve(ctx, &pb.SolveRequest{
		Year: year,
		Day:  day,
		Part: part,
	})

	return err
}

func loadExistingIssue(ctx context.Context, rsclient rstore_client.RStoreClient) (int32, error) {
	data, err := rsclient.Read(ctx, &rspb.ReadRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
	if err != nil {
		return -1, nil
	}
	return int32(binary.LittleEndian.Uint32(data.GetValue().GetValue())), nil
}

func raiseIssue(ctx context.Context, ghclient ghb_client.GithubridgeClient, rsclient rstore_client.RStoreClient, year, day, part int) error {
	issue, err := ghclient.AddIssue(ctx, &ghbpb.AddIssueRequest{Title: fmt.Sprintf("Solve %v - %v - %v", year, day, part), Job: "adventofcode"})
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
			err := solve(int32(time.Now().Year()), int32(day), int32(part))
			if status.Code(err) == codes.InvalidArgument {
				//Raise the issue to solve this problem
				return raiseIssue(ctx, ghclient, rsclient, year, day, part)
			}
		}
	}

	return nil
}

func main() {
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
		return
	}

	// If we're in a set, run this
	if time.Now().Month() == time.December && time.Now().Day() <= 25 {
		runYear(ctx, ghclient, rstore, time.Now().Year(), time.Now().Day())
		return
	}

	// If we're not in a set, work days at a time
	for day := 1; day <= 25; day++ {
		for year := 2015; year < time.Now().Year(); year++ {
			runYear(ctx, ghclient, rstore, year, day)
			return
		}
	}

	log.Printf("No more puizzles to solve")
}
