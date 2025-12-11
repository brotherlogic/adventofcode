package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day9TestInput = `
	7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
)

func TestDay9Part1(t *testing.T) {
	s := &Server{}
	res, err := s.Day9Part1(context.Background(), &pb.SolveRequest{
		Data: day9TestInput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 50 {
		t.Errorf("Expected 50, got %v", res.GetBigAnswer())
	}
}
func TestDay9Part2(t *testing.T) {
	s := &Server{}
	res, err := s.Day9Part2(context.Background(), &pb.SolveRequest{
		Data: day9TestInput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 24 {
		t.Errorf("Expected 24, got %v", res.GetBigAnswer())
	}
}

func TestRectangle(t *testing.T) {
	rect := getRectangle([]int64{11, 7}, []int64{7, 1})

	if rect != 35 {
		t.Errorf("Bad rectangle: %v", rect)
	}
}

func TestIntersection(t *testing.T) {

	i := lineIntersects([]int64{11, 1}, []int64{2, 1}, []int64{7, 3}, []int64{7, 1})
	if !i {
		t.Errorf("Line should intersect")
	}
}
