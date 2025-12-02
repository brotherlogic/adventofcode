package main

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func isInvalid(num int64) bool {
	strnum := strconv.Itoa(int(num))

	// Odd number of digits ; not invalid
	if len(strnum)%2 == 1 {
		return false
	}

	if strnum[0:len(strnum)/2] == strnum[len(strnum)/2:len(strnum)] {
		return true
	}
	return false
}

func (s *Server) Day2Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	total := int64(0)
	for _, r := range strings.Split(req.GetData(), ",") {
		pieces := strings.Split(strings.TrimSpace(r), "-")
		le, err := strconv.ParseInt(pieces[0], 10, 64)
		if err != nil {
			return nil, err
		}
		he, err := strconv.ParseInt(pieces[1], 10, 64)
		if err != nil {
			return nil, err
		}

		for i := le; i <= he; i++ {
			if isInvalid(i) {
				//log.Printf("Found %v", i)
				total += i
			}
		}
	}

	return &pb.SolveResponse{BigAnswer: total}, nil
}
