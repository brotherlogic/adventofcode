package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day10TestData = `
	[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`
	day10TestSO = `[#.....#.#] (0,1,2,3,4,6,7) (0,1,2,3,5,6,8) (6,8) (0,1,3,4,6,8) (0,2,3,5) (0,2,4,5,7,8) (0,1,4,6) (0,5,6,8) (0,1,2,3,5,8) {216,32,64,46,39,195,187,34,188}`
)

func TestDay10Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day10Part1(context.Background(), &pb.SolveRequest{
		Data: day10TestData,
	})

	if err != nil {
		t.Fatalf("Unable to run: %v", err)
	}

	if res.GetAnswer() != 7 {
		t.Errorf("Expected 7, got %v", res)
	}
}

func TestDay10Part1_Validate(t *testing.T) {
	s := &Server{}

	res, err := s.Day10Part1(context.Background(), &pb.SolveRequest{
		Data: day10TestSO,
	})

	if err != nil {
		t.Fatalf("Unable to run: %v", err)
	}

	if res.GetAnswer() != 8 {
		t.Errorf("Expected 8, got %v", res)
	}
}
