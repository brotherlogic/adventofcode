package main

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func findBiggest(num string, offset, cutoff int) (int64, int) {
	biggest := "0"[0]
	bo := 0
	for i := offset; i <= cutoff; i++ {
		if num[i] > biggest {
			biggest = num[i]
			bo = i
		}
	}

	v, _ := strconv.ParseInt(string(biggest), 10, 64)
	return v, bo
}

func (s *Server) Day3Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sum := int32(0)
	for _, line := range strings.Split(req.GetData(), "\n") {
		bv, bo := findBiggest(strings.TrimSpace(line), 0, len(line)-2)

		sv, _ := findBiggest(strings.TrimSpace(line), bo+1, len(line)-1)
		//log.Printf("Found %v and %v", bv, sv)
		sum += int32(bv*10 + sv)
	}

	return &pb.SolveResponse{Answer: sum}, nil
}
