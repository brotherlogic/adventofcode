package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

// Easier because these are all straightlines
func lineIntersects(ls, le, ps, pe []int) bool {
	if ls[0] == le[0] && ps[0] == pe[0] {
		// L is vertical, P is verticaal
		return false
	}

	if ls[1] == le[1] && ps[1] == pe[1] {
		// L is horizontal, P is vertical
		return false
	}

	if ls[0] == le[0] {
		// L is vertical, P is horizontal
		if ps[0] > pe[0] {
			if ls[0] < ps[0] && ls[0] > pe[0] {
				return true
			}
		}

		if ps[0] < pe[0] {
			if ls[0] > ps[0] && ls[0] < pe[0] {
				return true
			}
		}
	}

	if ls[1] == le[1] {
		// L is vertical, P is horizontal
		if ps[1] > pe[1] {
			if ls[1] < ps[1] && ls[1] > pe[1] {
				return true
			}
		}

		if ps[0] < pe[0] {
			if ls[1] > ps[1] && ls[1] < pe[1] {
				return true
			}
		}
	}

	return false
}

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
				log.Printf("Getting %v -> %v: %v / %v", tl, br, getRectangle(tl, br), getRectangle(br, tl))
				best = rect
			}
		}
	}

	return &pb.SolveResponse{BigAnswer: best}, nil
}
