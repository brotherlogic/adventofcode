package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day4 = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
)

func TestDay4Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day4Part1(context.Background(), &pb.SolveRequest{
		Data: day4,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 13 {
		t.Errorf("Expected 13, got %v", res.GetAnswer())
	}
}
