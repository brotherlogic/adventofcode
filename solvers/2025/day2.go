package main

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func isInvalid(num int64) bool {
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
				total += i
			}
		}
	}
}
