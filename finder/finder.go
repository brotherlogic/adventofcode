package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/adventofcode/proto"
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

	conn, err := grpc.Dial("adventofcode:8080", grpc.WithInsecure())
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

func loadExistingIssue() (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := grpc.Dial("githubridge:8080", grpc.WithInsecure())
	if err != nil {
		return -1, err
	}

	client := ghbpb.NewGithubBridgeClient(conn)

	return -1, nil
}

func raiseIssue(year, day, part int) error {
	return -1, nil
}

func findIssue(iid int32) error {
	return nil
}

func runYear(year int, db int) error {
	for day := 1; day <= db; day++ {
		for part := 1; part <= 2; part++ {
			err := solve(int32(time.Now().Year()), int32(day), int32(part))
			if status.Code(err) == codes.InvalidArgument {
				//Raise the issue to solve this problem
				return raiseIssue(year, day, part)
			}
		}
	}
}

func main() {
	// Check on the existing issue
	iid, err := loadExistingIssue()
	if err != nil {
		log.Fatalf("unable to load existing issue: %v", err)
	}

	// We have no solved the current issue
	if iid > 0 {
		return
	}

	// If we're in a set, run this
	if time.Now().Month() == time.December && time.Now().Day() <= 25 {
		runYear(time.Now().Year(), time.Now().Day())
		return
	}

	// If we're not in a set, work upwards
	for year := 2015; year < time.Now().Year(); year++ {
		runYear(year, 25)
		return
	}

	log.Printf("We've done it all")
}
