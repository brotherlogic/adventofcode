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
		if err == nil {
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
	if err == nil {
		if sol.GetSolution().GetBigAnswer() == res.GetBigAnswer() || sol.GetSolution().GetAnswer() == res.GetAnswer() {
			// We solved it
			return nil
		}

		return status.Errorf(codes.FailedPrecondition, "Solution is not present or incorrect %v vs %v", sol.GetSolution(), res)
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

func (f *finder) raiseIssue(ctx context.Context, year, day, part int32, err error) error {
	issue, err := f.ghclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{Title: fmt.Sprintf("Solve %v - %v - %v", year, day, part), Repo: "adventofcode", User: "brotherlogic"})
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

func (f *finder) runYear(ctx context.Context, ghclient ghb_client.GithubridgeClient, rsclient rstore_client.RStoreClient, year, db int32, issue *pb.Issue) error {
	for day := int32(db); day >= 1; day-- {
		for part := int32(1); part <= 2; part++ {
			if day == 25 && part == 2 {
				continue
			}
			err := f.solve(ctx, int32(year), int32(day), int32(part), issue)
			log.Printf("Solved %v %v %v -> %v", year, day, part, err)
			if status.Code(err) != codes.OK {
				if status.Code(err) == codes.Internal {
					// Something went wrong, silently fail this
					log.Printf("Processing error: %v", err)
					return nil
				}
				log.Printf("Raising issue for error: %v", err)
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

	if status.Code(err) == codes.NotFound {
		// Effective issue close
		log.Printf("Closing issue, as gh reported not found: %v", err)
		rissue = &ghbpb.GetIssueResponse{State: "closed"}
	} else if err != nil {
		return err
	}

	if rissue.GetState() == "closed" {
		log.Printf("Deleting issue - marked as clsoed")
		_, err := f.rsclient.Delete(ctx, &rspb.DeleteRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
		return err
	}

	// See if we have the right solution for this one
	conn, err := grpc.Dial("adventofcode.adventofcode:8082", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}
	connm, err := grpc.Dial("adventofcode.adventofcode:8080", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}
	client := pb.NewAdventOfCodeInternalServiceClient(conn)
	mclient := pb.NewAdventOfCodeServiceClient(connm)
	msol, err := mclient.Solve(ctx, &pb.SolveRequest{
		Year: issue.GetYear(),
		Day:  issue.GetDay(),
		Part: issue.GetPart(),
	})

	// If we haven't got a solution yet, we need to keep working
	if err != nil {
		log.Printf("Returning as we have no solution: %v", err)
		return nil
	}

	_, err = client.GetSolution(ctx, &pb.GetSolutionRequest{
		Year: issue.GetYear(),
		Day:  issue.GetDay(),
		Part: issue.GetPart(),
	})
	if status.Code(err) == codes.NotFound {
		// We have a potential solution, but no confirmation - if this is new, post
		found := false
		for _, sol := range issue.GetSolutionAttempts() {
			if sol.GetYear() == issue.GetYear() && sol.GetDay() == issue.GetDay() && sol.GetPart() == issue.GetPart() {
				if sol.GetAnswer() == msol.GetAnswer() &&
					sol.GetBigAnswer() == msol.GetBigAnswer() &&
					sol.GetStringAnswer() == msol.GetStringAnswer() {
					found = true
				}
			}
		}

		if found {
			return status.Errorf(codes.DataLoss, "Already seen this solution, no support for it or incorrect")
		}

		if !found {
			log.Printf("Found new solution: %v", msol)
			issue.SolutionAttempts = append(issue.SolutionAttempts, &pb.Solution{
				Answer:       msol.GetAnswer(),
				BigAnswer:    msol.GetBigAnswer(),
				StringAnswer: msol.GetStringAnswer(),
				Year:         issue.GetYear(),
				Day:          issue.GetDay(),
				Part:         issue.GetPart(),
			})
			data, err := proto.Marshal(issue)
			if err != nil {
				return err
			}
			_, err = f.rsclient.Write(ctx, &rspb.WriteRequest{Key: "brotherlogic/adventofcode/finder/cissue", Value: &anypb.Any{Value: data}})
			if err != nil {
				return err
			}

			_, err = f.ghclient.CommentOnIssue(ctx, &ghbpb.CommentOnIssueRequest{
				User:    "brotherlogic",
				Repo:    "adventofcode",
				Id:      int32(issue.GetId()),
				Comment: fmt.Sprintf("%v", msol),
			})
			return err
		}
	} else if status.Code(err) == codes.OK {
		// Close the issue
		_, err := f.ghclient.CloseIssue(ctx, &ghbpb.CloseIssueRequest{
			User: "brotherlogic",
			Repo: "adventofcode",
			Id:   int32(issue.GetId()),
		})
		return err
	}

	return nil
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func main() {
	log.Print("Running finder script")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
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
	log.Printf("Loaded existing issue: %v %v", issue, err)

	// We have no solved the current issue
	if issue != nil && issue.GetOpen() {
		log.Printf("Issue exists: %v", issue)
		log.Printf("Processed: %v", f.processNewIssue(ctx, issue))

		return
	}

	log.Printf("Running for the current year")

	loc, _ := time.LoadLocation("America/New_York")
	ctime := time.Now().In(loc)

	// If we're in a set, run this
	if ctime.Month() == time.December {
		log.Printf("In a set: %v", ctime.Day())
		err = f.runYear(ctx, ghclient, rstore, int32(ctime.Year()), min(25, int32(ctime.Day())), issue)
		if err != nil {
			log.Printf("Result: %v", err)
			return
		}

		return
	}

	log.Println("Not running for the current year - trying other years")

	// If we're not in a set, work days at a time.
	for day := int32(1); day <= 25; day++ {
		for year := 2015; year < time.Now().Year(); year++ {
			if f.runYear(ctx, ghclient, rstore, int32(year), day, issue) != nil {
				return
			}
		}
	}

	log.Printf("No more puizzles to solve")
}
