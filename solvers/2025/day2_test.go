package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	TEST = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
)

func TestDay2Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day2Part1(context.Background(), &pb.SolveRequest{
		Data: TEST,
	})

	if err != nil {
		t.Fatalf("Unable to run test: %v", err)
	}

	if res.GetBigAnswer() != 1227775554 {
		t.Errorf("Got %v should have been 1227775554", res.GetBigAnswer())
	}
}

func TestDay2Part2(t *testing.T) {
	s := &Server{}

	res, err := s.Day2Part2(context.Background(), &pb.SolveRequest{
		Data: TEST,
	})

	if err != nil {
		t.Fatalf("Unable to run test: %v", err)
	}

	if res.GetBigAnswer() != 4174379265 {
		t.Errorf("Got %v should have been 4174379265", res.GetBigAnswer())
	}
}
