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
	day11testinput2 = `
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
)

func TestDay11Part1(t *testing.T) {
	s := &Server{}

	res, err := s.Day11Part1(context.Background(), &pb.SolveRequest{
		Data: day11testinput,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 5 {
		t.Errorf("Expected 5, got %v", res.GetAnswer())
	}
}

func TestDay11Part2(t *testing.T) {
	s := &Server{}

	res, err := s.Day11Part2(context.Background(), &pb.SolveRequest{
		Data: day11testinput2,
	})

	if err != nil {
		t.Fatalf("Unable to solve: %v", err)
	}

	if res.GetBigAnswer() != 2 {
		t.Errorf("Expected 2, got %v", res.GetAnswer())
	}
}

func TestAvoidLooping(t *testing.T) {
	s := &Server{}

	c, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := s.Day11Part1(c, &pb.SolveRequest{
		Data: `you: b
		b: you`,
	})
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}

	if res.GetAnswer() != 0 {
		t.Errorf("Bad answer: %v", res)
	}
}
