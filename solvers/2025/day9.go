package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

// Easier because these are all straightlines
func lineIntersects(ls, le, ps, pe []int64) bool {
	if ls[0] == le[0] && ps[0] == pe[0] {
		// L is vertical, P is verticaal
		return false
	}

	if ls[1] == le[1] && ps[1] == pe[1] {
		// L is horizontal, P is horizontal
		return false
	}

	if ls[0] == le[0] {
		// L is vertical, P is horizontal
		cx, cy := ls[0], ps[1]

		xin := false
		yin := false
		if ps[0] > pe[0] && (cx <= ps[0] && cx >= pe[0]) {
			xin = true
		}
		if ps[0] < pe[0] && (cx >= ps[0] && cx <= pe[0]) {
			xin = true
		}
		if ls[1] > le[1] && (cy <= ls[1] && cy >= le[1]) {
			yin = true
		}
		if ls[1] < le[1] && (cy >= ls[1] && cy <= le[1]) {
			yin = true
		}

		if xin && yin {
			// Not an intersection if one of the crosspoints is at a corner
			if (cx == ls[0] && cy == ls[1]) || (cx == le[0] && cy == le[1]) {
				return false
			}

			return true
		}
	}

	if ls[1] == le[1] {
		// L is horiztonal, P is vertical
		cx, cy := ps[0], ls[1]

		xin := false
		yin := false
		if ls[0] > le[0] && (cx <= ls[0] && cx >= le[0]) {
			xin = true
		}
		if ls[0] < le[0] && (cx >= ls[0] && cx <= le[0]) {
			xin = true
		}
		if ps[1] > pe[1] && (cy <= ps[1] && cy >= pe[1]) {
			yin = true
		}
		if ps[1] < pe[1] && (cy >= ps[1] && cy <= pe[1]) {
			yin = true
		}

		if xin && yin {
			// Not an intersection if one of the crosspoints is at a corner
			if (cx == ls[0] && cy == ls[1]) || (cx == le[0] && cy == le[1]) {
				return false
			}

			return true
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
				best = rect
			}
		}
	}

	return &pb.SolveResponse{BigAnswer: best}, nil
}

func intersect(x1, y1, x2, y2 int64, coords [][]int64) bool {
	for i := 0; i < len(coords)-1; i++ {
		if lineIntersects([]int64{x1, y1}, []int64{x2, y2}, coords[i], coords[i+1]) {
			return true
		}
	}

	return lineIntersects([]int64{x1, y1}, []int64{x2, y2}, coords[0], coords[len(coords)-1])

}

func rectangleIntersects(tl, br []int64, coords [][]int64) bool {
	if intersect(tl[0], tl[1], tl[0], br[1], coords) {
		return true
	}

	if intersect(tl[0], br[1], br[0], br[1], coords) {
		return true
	}

	if intersect(br[0], br[1], br[0], tl[1], coords) {
		return true
	}

	if intersect(br[0], tl[1], tl[0], tl[1], coords) {

		return true
	}
	return false
}

func (s *Server) Day9Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	best := int64(0)

	coords := buildCoords(req.GetData())

	for i, tl := range coords {
		for _, br := range coords[i+1:] {
			rect := getRectangle(tl, br)

			intersects := rectangleIntersects(tl, br, coords)

			if !intersects && rect > best {
				best = rect
			}
		}
	}

	return &pb.SolveResponse{BigAnswer: best}, nil
}
