package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func TestDay3Part1(t *testing.T) {
	banks := `987654321111111
811111111111119
234234234234278
818181911112111`

	s := Server{}
	res, err := s.Day3Part1(context.Background(), &pb.SolveRequest{
		Data: banks,
	})
	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 357 {
		t.Errorf("Expected 357, got %v", res.GetAnswer())
	}
}
