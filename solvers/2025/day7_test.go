package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day7TestInput = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

	day7TestSimple = `..S..
..^..
.^.^.
.....`
)

func TestDay7Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day7Part1(context.Background(), &pb.SolveRequest{
		Data: day7TestInput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 21 {
		t.Errorf("Expected 21, got %v", res.GetAnswer())
	}
}

func TestDay7Part2(t *testing.T) {
	s := &Server{}

	res, err := s.Day7Part2(context.Background(), &pb.SolveRequest{
		Data: day7TestInput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 40 {
		t.Errorf("Expected 40, got %v", res.GetAnswer())
	}
}

func TestDay7Part2_Simple(t *testing.T) {
	s := &Server{}

	res, err := s.Day7Part2(context.Background(), &pb.SolveRequest{
		Data: day7TestSimple,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 4 {
		t.Errorf("Expected 2, got %v", res.GetAnswer())
	}
}
