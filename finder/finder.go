package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	ghb_client "github.com/brotherlogic/githubridge/client"
	pstore_client "github.com/brotherlogic/pstore/client"

	pb "github.com/brotherlogic/adventofcode/proto"
	ghbpb "github.com/brotherlogic/githubridge/proto"
	pspb "github.com/brotherlogic/pstore/proto"
)

var (
	solvingDuration = time.Minute * 5
	retries         = 3
)

type finder struct {
	ghclient ghb_client.GithubridgeClient
	psclient pstore_client.PStoreClient
}

func download(year, day int32, c string) (string, error) {
	client := new(http.Client)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "github.com/brotherlogic/advent-of-code-finder")

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", c
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Printf("Respose: %v", string(b))
		return "", errors.New(resp.Status)
	}

	return (string(b)), nil
}

func (f *finder) getData(ctx context.Context, year, day int32) error {
	log.Printf("Retrieving data for %v / %v", year, day)

	cookie, err := f.psclient.Read(ctx, &pspb.ReadRequest{Key: "brotherlogic/adventofcode/finder/cookie"})
	if err != nil {
		return err
	}

	data, err := download(year, day, string(cookie.GetValue().GetValue()))
	if err != nil {
		return err
	}

	connm, err := grpc.Dial("adventofcode.adventofcode:8082", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}
	client := pb.NewAdventOfCodeInternalServiceClient(connm)
	_, err = client.Upload(ctx, &pb.UploadRequest{Day: day, Year: year, Data: data, Part: 1})
	log.Printf("Uploaded: %v -> %v", data, err)
	return err
}

func (f *finder) addLabel(ctx context.Context, label string, issue *pb.Issue) error {
	_, err := f.ghclient.AddLabel(ctx, &ghbpb.AddLabelRequest{
		User:  "brotherlogic",
		Repo:  "adventofcode",
		Id:    int32(issue.GetId()),
		Label: label,
	})
	return err
}

func (f *finder) hasLabel(ctx context.Context, label string, issue *pb.Issue) (bool, error) {
	labels, err := f.ghclient.GetLabels(ctx, &ghbpb.GetLabelsRequest{
		Id:   int32(issue.GetId()),
		Repo: "adventofcode",
		User: "brotherlogic",
	})
	if err != nil {
		return false, err
	}

	for _, flabel := range labels.GetLabels() {
		if label == flabel {
			return true, nil
		}
	}
	return false, nil
}

func (f *finder) removeLabel(ctx context.Context, label string, issue *pb.Issue) error {
	_, err := f.ghclient.DeleteLabel(ctx, &ghbpb.DeleteLabelRequest{
		User:  "brotherlogic",
		Repo:  "adventofcode",
		Id:    int32(issue.GetId()),
		Label: label,
	})
	return err
}

func (f *finder) solve(ctx context.Context, year, day, part int32, issue *pb.Issue) error {

	log.Printf("Solving %v %v %v", year, day, part)
	for i := 0; i < retries; i++ {
		err := f.solveInternal(ctx, year, day, part, issue)
		if err == nil {
			return err
		}

		if status.Code(err) == codes.Unimplemented {
			f.addLabel(ctx, "Requires Implementation", issue)
		}

		log.Printf("Solve fail: %v", err)
	}

	return status.Errorf(codes.ResourceExhausted, "Unable to solve with retries")
}

func (f *finder) validateYear(ctx context.Context, year int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), solvingDuration)
	defer cancel()

	conn, err := grpc.Dial("adventofcode.adventofcode:8080", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}
	defer conn.Close()

	client := pb.NewAdventOfCodeServiceClient(conn)
	_, err = client.Solve(ctx, &pb.SolveRequest{Year: year})
	return err
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
	data, err := f.psclient.Read(ctx, &pspb.ReadRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
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
	var issue *ghbpb.CreateIssueResponse
	var ierr error
	if status.Code(err) == codes.InvalidArgument {
		issue, ierr = f.ghclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{
			Title: fmt.Sprintf("Add infrastructure for %v", year),
			Repo:  "adventofcode",
			User:  "brotherlogic",
		})
		if ierr != nil {
			return ierr
		}
	} else if status.Code(err) == codes.FailedPrecondition {
		issue, ierr = f.ghclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{Title: fmt.Sprintf("Add zero day testbed for %v", year), Repo: "adventofcode", User: "brotherlogic"})
		if ierr != nil {
			return ierr
		}
	} else {
		issue, ierr = f.ghclient.CreateIssue(ctx, &ghbpb.CreateIssueRequest{Title: fmt.Sprintf("Solve %v - %v - %v", year, day, part), Repo: "adventofcode", User: "brotherlogic"})
		if ierr != nil {
			return ierr
		}
	}

	iss := &pb.Issue{
		Id:            issue.GetIssueId(),
		Open:          true,
		Year:          year,
		Day:           day,
		Part:          part,
		LastErrorCode: fmt.Sprintf("%v", status.Code(err)),
	}
	bytes, err := proto.Marshal(iss)
	if err != nil {
		return err
	}

	_, err = f.psclient.Write(ctx, &pspb.WriteRequest{Key: "brotherlogic/adventofcode/finder/cissue", Value: &anypb.Any{Value: bytes}})
	log.Printf("Written issue: %v", err)
	return err
}

func (f *finder) runYear(ctx context.Context, ghclient ghb_client.GithubridgeClient, psclient pstore_client.PStoreClient, year, db int32, issue *pb.Issue) error {
	log.Printf("Running for year %v with day %v", year, db)
	if db == 0 {
		err := f.validateYear(ctx, year)
		log.Printf("Validated %v -> %v", year, err)
		if err != nil {
			err2 := f.raiseIssue(ctx, year, 0, 0, err)
			if err2 != nil {
				return err2
			}
		}
		return err
	}

	conn, err := grpc.Dial("adventofcode.adventofcode:8082", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial aoc: %w", err)
	}
	client := pb.NewAdventOfCodeInternalServiceClient(conn)
	defer conn.Close()

	for day := int32(db); day >= 1; day-- {
		for part := int32(1); part <= 2; part++ {
			if day == 25 && part == 2 {
				continue
			}

			// Look to see if we already have a solution for this
			_, err = client.GetSolution(ctx, &pb.GetSolutionRequest{
				Year: year,
				Day:  day,
				Part: part,
			})
			if status.Code(err) == codes.OK {
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
				log.Printf("Raising issue for this error: %v", err)
				//Raise the issue to solve this problem
				err2 := f.raiseIssue(ctx, year, day, part, err)
				log.Printf("Issue %v for %v", err2, err)
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
		_, err := f.psclient.Delete(ctx, &pspb.DeleteRequest{Key: "brotherlogic/adventofcode/finder/cissue"})
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

	// This means we can't find the data to run the solution
	log.Printf("Solve error: %v", err)
	if status.Code(err) == codes.NotFound {
		err = f.addLabel(ctx, "Needs Data", issue)
		log.Printf("Added needs data label: %v", err)
		err = f.getData(ctx, issue.GetYear(), issue.GetDay())
		if err != nil {
			log.Printf("Data miss: %v", err)
			f.addLabel(ctx, "Data Issue", issue)

			if status.Code(err) == codes.NotFound {
				f.addLabel(ctx, "Cookie Missing", issue)
			}
		}
	} else {
		f.removeLabel(ctx, "Needs Data", issue)
		f.removeLabel(ctx, "Data Issue", issue)
		f.removeLabel(ctx, "Cookie Missing", issue)
	}

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

	if issue.GetLastErrorCode() == "INVALID_ARGUMENT" && status.Code(err) != codes.InvalidArgument {
		// Close the issue
		_, err := f.ghclient.CloseIssue(ctx, &ghbpb.CloseIssueRequest{
			User: "brotherlogic",
			Repo: "adventofcode",
			Id:   int64(issue.GetId()),
		})
		return err
	}

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
			// See if this solution is correct
			labe, err := f.hasLabel(ctx, "correct", issue)
			if err != nil {
				log.Fatalf("Bad label: %v", err)
			}

			if labe {
				bsol := issue.GetSolutionAttempts()[0]
				for _, sol := range issue.GetSolutionAttempts() {
					if sol.GetSolutionMade() > bsol.GetSolutionMade() {
						bsol = sol
					}
				}

				_, err = client.AddSolution(ctx, &pb.AddSolutionRequest{Solution: bsol})
				log.Printf("Added solution: %v (%v) from %v", err, bsol, issue.GetSolutionAttempts())
			}

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
				SolutionMade: time.Now().Unix(),
			})
			data, err := proto.Marshal(issue)
			if err != nil {
				return err
			}
			_, err = f.psclient.Write(ctx, &pspb.WriteRequest{Key: "brotherlogic/adventofcode/finder/cissue", Value: &anypb.Any{Value: data}})
			if err != nil {
				return err
			}

			if msol.GetAnswer() > 0 || msol.GetBigAnswer() > 0 || msol.GetStringAnswer() != "" {
				_, err = f.ghclient.CommentOnIssue(ctx, &ghbpb.CommentOnIssueRequest{
					User:    "brotherlogic",
					Repo:    "adventofcode",
					Id:      int32(issue.GetId()),
					Comment: fmt.Sprintf("Solution: %v", msol),
				})
				return err
			}
			return nil
		}
	} else if status.Code(err) == codes.OK {
		// Close the issue
		_, err := f.ghclient.CloseIssue(ctx, &ghbpb.CloseIssueRequest{
			User: "brotherlogic",
			Repo: "adventofcode",
			Id:   int64(issue.GetId()),
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

	ghclient, err := ghb_client.GetClientInternal()
	if err != nil {
		log.Fatalf("unable to get ghb client: %v", err)
	}

	pstore, err := pstore_client.GetClient()
	if err != nil {
		log.Fatalf("unable to get pstore client: %v", err)
	}

	f := &finder{
		ghclient: ghclient,
		psclient: pstore,
	}

	// Check on the existing issue
	issue, err := f.loadExistingIssue(ctx)
	if err != nil && status.Code(err) != codes.NotFound {
		log.Fatalf("unable to load existing issue: %v", err)
	}
	log.Printf("Loaded existing issue: %v %v", issue, err)

	// We have not solved the current issue
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
		err = f.runYear(ctx, ghclient, pstore, int32(ctime.Year()), min(25, int32(ctime.Day())), issue)
		if err != nil {
			log.Printf("Result: %v", err)
			return
		}

		return
	}

	log.Println("Not running for the current year - trying other years")

	// If we're not in a set, work days at a time.
	// First check that we have all the infrastructure we need

	for year := ctime.Year(); year >= 2015; year-- {
		err = f.runYear(ctx, ghclient, pstore, int32(year), 0, issue)
		if err != nil {
			log.Printf("Result: %v", err)
			return
		}
	}

	log.Printf("No more puizzles to solve")
}
