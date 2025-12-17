package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day12TestInput = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`
)

func TestDay12Part1(t *testing.T) {
	s := &Server{}
	res, err := s.Day12Part1(context.Background(), &pb.SolveRequest{
		Data: day12TestInput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 2 {
		t.Errorf("Expected 2, got %v", res.GetAnswer())
	}
}
