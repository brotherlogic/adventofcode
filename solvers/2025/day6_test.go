package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day6 = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`
)

func TestDay6Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day6Part1(context.Background(), &pb.SolveRequest{
		Data: day6,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 4277556 {
		t.Errorf("Expected 4277556, got %v", res.GetAnswer())
	}
}
