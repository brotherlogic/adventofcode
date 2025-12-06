package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day5 = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
)

func TestDay5Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day5Part1(context.Background(), &pb.SolveRequest{
		Data: day5,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 3 {
		t.Errorf("Expected 3, got %v", res.GetAnswer())
	}
}

func TestDay5Part2(t *testing.T) {
	s := &Server{}

	res, err := s.Day5Part2(context.Background(), &pb.SolveRequest{
		Data: day5,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 14 {
		t.Errorf("Expected 14, got %v", res.GetBigAnswer())
	}
}
