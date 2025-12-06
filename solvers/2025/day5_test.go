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

func TestCollapseComplex(t *testing.T) {
	var nums [][]int64
	nums = append(nums, []int64{6, 20})
	nums = append(nums, []int64{23, 200})
	nums = append(nums, []int64{15, 25})

	c := metaCollapse(nums)
	if len(c) != 1 {
		t.Errorf("Miscollapse of broken ranges: %v", c)
	}
	if c[0][0] != int64(6) || c[0][1] != int64(200) {
		t.Errorf("Miscollapse of broken ranges: %v", c)
	}
}

func TestCollapseComplex2(t *testing.T) {
	s := &Server{}

	res, err := s.Day5Part2(context.Background(), &pb.SolveRequest{
		Data: `1-10
2-5`,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 10 {
		t.Errorf("Expected 15, got %v", res.GetBigAnswer())
	}

}
