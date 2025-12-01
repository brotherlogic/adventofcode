package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func TestDay1Part1(t *testing.T) {
	data := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	server := &Server{}
	resp, err := server.Day1Part1(context.Background(), &pb.SolveRequest{Data: data})
	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}
	if resp.GetAnswer() != 3 {
		t.Fatalf("Bad answer: %v", resp.GetAnswer())
	}
}

func TestDay1Part2(t *testing.T) {
	data := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	server := &Server{}
	resp, err := server.Day1Part2(context.Background(), &pb.SolveRequest{Data: data})
	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}
	if resp.GetAnswer() != 6 {
		t.Fatalf("Bad answer: %v", resp.GetAnswer())
	}
}
