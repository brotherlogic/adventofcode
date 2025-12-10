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

func TestRectangle(t *testing.T) {
	rect := getRectangle([]int64{11, 7}, []int64{7, 1})

	if rect != 35 {
		t.Errorf("Bad rectangle: %v", rect)
	}
}

func TestIntersection(t *testing.T) {
	i := lineIntersects([]int{2, 2}, []int{10, 2}, []int{5, 5}, []int{5, 2})
	if !i {
		t.Errorf("Line should intersect")
	}

	i = lineIntersects([]int{5, 2}, []int{10, 2}, []int{5, 5}, []int{5, 2})
	if i {
		t.Errorf("Line should not intersect")
	}
}
