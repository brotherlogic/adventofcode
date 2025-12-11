package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day11testinput = `
	aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
)

func TestDay11Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day11Part1(context.Background(), &pb.SolveRequest{
		Data: day11testinput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 5 {
		t.Errorf("Expected 5, got %v", res.GetBigAnswer())
	}
}

func TestAvoidLooping(t *testing.T) {
	s := &Server{}

	c, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	res, err := s.Day11Part1(c, &pb.SolveRequest{
		Data: `a: b
		b:a`,
	})
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}

	if res.GetAnswer() != 0 {
		t.Errorf("Bad answer: %v", res)
	}
}
