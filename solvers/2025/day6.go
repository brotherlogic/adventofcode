package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func buildOffsetGrid(grid string) ([][]string, []string) {
	elems := strings.Split(strings.TrimSpace(grid), "\n")

	opsLine := elems[len(elems)-1]
	var ops []string
	var lineOffsets []int
	for offset, ch := range opsLine {
		if ch == '*' || ch == '+' {
			lineOffsets = append(lineOffsets, offset)
			ops = append(ops, string(ch))
		}
	}
	mlen := 0
	for _, line := range elems {
		if len(line) > mlen {
			mlen = len(line)
		}
	}

	// Append off the length of the line also
	lineOffsets = append(lineOffsets, mlen)

	var data [][]string
	for _, line := range elems[:len(elems)-1] {
		var dline []string
		for i := 0; i < len(lineOffsets)-1; i++ {
			dline = append(dline, line[lineOffsets[i]:lineOffsets[i+1]])
		}
		data = append(data, dline)
	}

	return data, ops
}

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

func resolve(numv []int64, op string) int64 {
	value := int64(0)

	biggest := int64(0)
	for _, num := range numv {
		if num > biggest {
			biggest = num
		}
	}

	topv := len(fmt.Sprintf("%v", biggest))

	for i := topv; i > 0; i-- {

	}

	return value
}

func resolveAdvanced(numv []string, op string) int64 {
	sumv := int64(0)
	if op == "*" {
		sumv = 1
	}

	for offset := 0; offset < len(numv[0]); offset++ {
		nnum := ""
		for _, numvv := range numv {
			if numvv[offset] != ' ' {
				nnum += string(numvv[offset])
			}
		}
		if len(nnum) > 0 {
			val, err := strconv.ParseInt(nnum, 10, 64)
			if err != nil {
				log.Fatalf("Cannot parse: %v", err)
			}

			if op == "*" {
				sumv *= val
			} else {
				sumv += val
			}
		}
	}
	return sumv
}

func (s *Server) Day6Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int64(0)

	data, ops := buildOffsetGrid(req.GetData())

	for i, op := range ops {
		var numv []string
		for _, parts := range data {
			numv = append(numv, parts[i])
		}

		sumv += resolveAdvanced(numv, op)
	}

	return &pb.SolveResponse{
		BigAnswer: sumv,
	}, nil
}
