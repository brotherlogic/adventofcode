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

type finder struct {
	ghclient ghb_client.GithubridgeClient
	rsclient rstore_client.RStoreClient
}

func (f *finder) solve(ctx context.Context, year, day, part int32, issue *pb.Issue) error {

	log.Printf("Solving %v %v %v", year, day, part)
	for i := 0; i < retries; i++ {
		err := f.solveInternal(ctx, year, day, part, issue)
		if status.Code(err) != codes.NotFound {
			return err
		}
		log.Printf("Solve fail: %v", err)
	}

	return status.Errorf(codes.ResourceExhausted, "Unable to solve with retries")
}

func (f *finder) solveInternal(sctx context.Context, year, day, part int32, issue *pb.Issue) error {
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

	if issue == nil {
		return f.raiseIssue(sctx, year, day, part, fmt.Errorf("Starting issue"))
	}

	if status.Code(err) == codes.NotFound {
		return addSolutionToIssue(ctx, sol.GetSolution(), issue)
	}

	return err
}

func (f *finder) loadExistingIssue(ctx context.Context) (*pb.Issue, error) {
	data, err := f.rsclient.Read(ctx, &rspb.ReadRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
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

func (f *finder) raiseIssue(ctx context.Context, year, day, part int32, err error) error {
	issue, err := f.ghclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{Title: fmt.Sprintf("Solve %v - %v - %v (%v)", year, day, part, err), Repo: "adventofcode", User: "brotherlogic"})
	if err != nil {
		return err
	}

	iss := &pb.Issue{
		Id:   issue.GetIssueId(),
		Open: true,
		Year: year,
		Day:  day,
		Part: part,
	}
	bytes, err := proto.Marshal(iss)
	if err != nil {
		return err
	}

	_, err = f.rsclient.Write(ctx, &rspb.WriteRequest{Key: "brotherlogic/adventofcode/finder/cissue", Value: &anypb.Any{Value: bytes}})
	log.Printf("Written issue: %v", err)
	return err
}

func findIssue(iid int32) error {
	return nil
}

func (f *finder) runYear(ctx context.Context, ghclient ghb_client.GithubridgeClient, rsclient rstore_client.RStoreClient, year, db int32, issue *pb.Issue) error {
	for day := int32(1); day <= db; day++ {
		for part := int32(1); part <= 2; part++ {
			err := f.solve(ctx, int32(year), int32(day), int32(part), issue)
			log.Printf("Solved %v %v %v -> %v", year, day, part, err)
			if status.Code(err) != codes.OK {
				//Raise the issue to solve this problem
				err2 := f.raiseIssue(ctx, year, day, part, err)
				if err2 != nil {
					return err2
				}
				return err
			}
		}
	}

	return nil
}

func (f *finder) processNewIssue(ctx context.Context, issue *pb.Issue) error {
	rissue, err := f.ghclient.GetIssue(ctx, &ghbpb.GetIssueRequest{
		User: "brotherlogic",
		Repo: "adventofcode",
		Id:   int32(issue.GetId()),
	})
	if err != nil {
		return err
	}

	if rissue.GetState() == "closed" {
		f.rsclient.Delete(ctx, &rspb.DeleteRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
		return nil
	}

	// See if we have the right solution for this one
	conn, err := grpc.Dial("adventofcode.adventofcode:8082", grpc.WithInsecure())
	connm, err := grpc.Dial("adventofcode.adventofcode:8080", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}

	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}
	client := pb.NewAdventOfCodeInternalServiceClient(conn)
	mclient := pb.NewAdventOfCodeServiceClient(connm)
	msol, err := mclient.Solve(ctx, &pb.SolveRequest{
		Year: rissue.GetYear(),
		Day:  rissue.GetDay(),
		Part: rissue.GetPart(),
	})

	// If we haven't got a solution yet, we need to keep working
	if err != nil {
		return nil
	}

	solution, err := client.GetSolution(ctx, &pb.GetSolutionRequest{
		Year: rissue.GetYear(),
		Day:  rissue.GetDay(),
		Part: rissue.GetPart(),
	})
	if status.Code(err) == codes.NotFound {
		// We have a potential solution, but no confirmation - if this is new, post
		found := false
		for _, sol := range rissue.GetSolutions() {
			if sol.GetYear() == rissue.GetYear() && sol.GetDay() == rissue.GetDay() && sol.GetPart() == rissue.GetPart() {
				if sol.GetAnswer() == solution.GetSolution().GetAnswer() &&
					sol.GetBigAnswer() == solution.GetSolution().GetBigAnswer() &&
					sol.GetStringAnswer() == solution.GetSolution().GetStringAnswer() {
					found = true
				}
			}
		}

		if !found {
			rissue.Solutions = append(rissue.Solutions, solution.GetSolution)
			data, err := proto.Marshal(rissue)
			if err != nil {
				return err
			}
			_, err = f.rsclient.Write(ctx, &rspb.WriteRequest{Key: "brotherlogic/adventofcode/finder/cissue", Value: &anypb.Any{Value: data}})
			if err != nil {
				return err
			}

			_, err := f.ghclient.Add
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

	f := &finder{
		ghclient: ghclient,
		rsclient: rstore,
	}

	// Check on the existing issue
	issue, err := f.loadExistingIssue(ctx)
	if err != nil && status.Code(err) != codes.NotFound {
		log.Fatalf("unable to load existing issue: %v", err)
	}

	// We have no solved the current issue
	if issue != nil && issue.GetOpen() {
		log.Printf("Issue exists: %v", issue)
		f.processNewIssue(ctx, issue)

		return
	}

	// If we're in a set, run this
	if time.Now().Month() == time.December && time.Now().Day() <= 25 {
		for day := int32(1); day <= int32(25); day++ {
			err = f.runYear(ctx, ghclient, rstore, int32(time.Now().Year()), day, issue)
			if err != nil {
				log.Printf("Result: %v", err)
				return
			}
		}
	}

	// If we're not in a set, work days at a time
	for day := int32(1); day <= 25; day++ {
		for year := 2015; year < time.Now().Year(); year++ {
			if f.runYear(ctx, ghclient, rstore, int32(year), day, issue) != nil {
				return
			}
		}
	}

	log.Printf("No more puizzles to solve")
}
