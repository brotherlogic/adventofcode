package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func buildGrid(grid string) ([][]int64, []string) {
	var data [][]int64
	var ops []string

	for _, line := range strings.Split(grid, "\n") {
		if strings.Contains(line, "*") {
			ops = strings.Fields(line)
		} else if len(line) > 0 {
			var numline []int64
			for _, num := range strings.Fields(line) {
				val, err := strconv.ParseInt(num, 10, 64)
				if err != nil {
					log.Fatalf("Cannot parse line: %v -> %v", line, err)
				}
				numline = append(numline, val)
			}
			data = append(data, numline)
		}
	}

	return data, ops
}

func (s *Server) Day6Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int64(0)

	data, ops := buildGrid(req.GetData())

	for i, op := range ops {
		val := int64(0)
		if op == "*" {
			val = 1
		}

		for _, parts := range data {
			if op == "*" {
				val *= parts[i]
			} else {
				val += parts[i]
			}
		}

		sumv += val
	}

	return &pb.SolveResponse{
		BigAnswer: sumv,
	}, nil
}

func (s *Server) Day6Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int64(0)

	data, ops := buildGrid(req.GetData())

	for i, op := range ops {
		val := int64(0)
		if op == "*" {
			val = 1
		}

		for _, parts := range data {
			if op == "*" {
				val *= parts[i]
			} else {
				val += parts[i]
			}
		}

		sumv += val
	}

	return &pb.SolveResponse{
		BigAnswer: sumv,
	}, nil
}
