package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
)

const (
	day8TestInput = `
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`
)

func TestDay8Part1(t *testing.T) {
	res, err := runDay8Part1(&pb.SolveRequest{
		Data: day8TestInput,
	}, 10)

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 40 {
		t.Errorf("Expected 40, got %v", res.GetAnswer())
	}
}

func TestDay8Part2(t *testing.T) {
	s := &Server{}
	res, err := s.Day8Part2(context.Background(), &pb.SolveRequest{
		Data: day8TestInput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetAnswer() != 25272 {
		t.Errorf("Expected 25272, got %v", res.GetAnswer())
	}
}
