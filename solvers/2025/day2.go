package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func isInvalid(num int64) bool {
	strnum := strconv.FormatInt(num, 10)

	// Odd number of digits ; not invalid
	if len(strnum)%2 == 1 {
		return false
	}

	if strnum[0:len(strnum)/2] == strnum[len(strnum)/2:len(strnum)] {
		return true
	}
	return false
}

func isInvalidFull(num int64) bool {
	strnum := strconv.FormatInt(num, 10)

	// Odd number of digits ; not invalid
	if len(strnum)%2 == 1 {
		return false
	}

	for rlen := 1; rlen <= len(strnum)/2; rlen++ {
		found := true
		for start := rlen; start <= len(strnum)-rlen; start++ {
			if strnum[0:rlen] != strnum[start:start+rlen] {
				found = false
				break
			}
		}
		if found {
			return true
		}
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

		log.Printf("Checking from %v to %v (%v)", le, he, he-le)
		for i := le; i <= he; i++ {
			if isInvalid(i) {
				//log.Printf("Found %v", i)
				total += i
			}
		}
	}

	return &pb.SolveResponse{BigAnswer: total}, nil
}

func (s *Server) Day2Part2(tx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
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

		log.Printf("Checking from %v to %v (%v)", le, he, he-le)
		for i := le; i <= he; i++ {
			if isInvalidFull(i) {
				log.Printf("Found %v", i)
				total += i
			}
		}
	}

	return &pb.SolveResponse{BigAnswer: total}, nil
}
