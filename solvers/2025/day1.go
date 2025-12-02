package main

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func (s *Server) Day1Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	position := int64(50)
	count := 0

	for _, line := range strings.Split(req.GetData(), "\n") {
		if len(line) > 0 {
			num, err := strconv.ParseInt(line[1:], 10, 64)
			if err != nil {
				return nil, err
			}
			if line[0] == 'L' {
				position -= num
				for position < 0 {
					position += 100
				}
			} else {
				position += num
				for position >= 100 {
					position -= 100
				}
			}
			if position == 0 {
				count++
			}
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}

func (s *Server) Day1Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	position := int64(50)
	count := 0

	for _, line := range strings.Split(req.GetData(), "\n") {
		if len(line) > 0 {
			num, err := strconv.ParseInt(line[1:], 10, 64)
			if err != nil {
				return nil, err
			}
			if line[0] == 'L' {
				// Knock one off
				if position == 0 {
					count--
				}
				position -= num
				for position < 0 {
					position += 100
					count++
				}
			} else {
				position += num
				for position > 100 {
					position -= 100
					count++
				}
			}
			if position == 0 || position == 100 {
				count++
				position = 0
			}
		}
	}

	return &pb.SolveResponse{Answer: int32(count)}, nil
}
