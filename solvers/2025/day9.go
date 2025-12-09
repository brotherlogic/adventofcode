package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func buildCoords(data string) [][]int64 {
	var coords [][]int64

	for _, line := range strings.Split(data, "\n") {
		if strings.Contains(line, ",") {
			pieces := strings.Split(strings.TrimSpace(line), ",")
			num1, err := strconv.ParseInt(pieces[0], 10, 32)
			if err != nil {
				log.Fatalf("Cannot parse: %v (%v)", err, line)
			}

			num2, err := strconv.ParseInt(pieces[1], 10, 32)
			if err != nil {
				log.Fatalf("Cannot parse: %v (%v)", err, line)
			}

			coords = append(coords, []int64{num1, num2})
		}
	}

	return coords
}

func abs(val int64) int64 {
	if val < 0 {
		return -val
	}
	return val
}

func getRectangle(x, y []int64) int64 {
	return (abs(x[0]-y[0]) + 1) * (abs(x[1]-y[1]) + 1)
}

func (s *Server) Day9Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	best := int64(0)

	coords := buildCoords(req.GetData())
	for i, tl := range coords {
		for _, br := range coords[i+1:] {
			rect := getRectangle(tl, br)
			if rect > best {
				best = rect
			}
		}
	}

	return &pb.SolveResponse{Answer: int32(best)}, nil
}
