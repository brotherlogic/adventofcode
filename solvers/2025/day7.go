package main

import (
	"context"
	"log"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func buildSplitGrid(data string) [][]int {
	var grid [][]int

	for _, line := range strings.Split(data, "\n") {
		var row []int
		for _, char := range line {
			switch char {
			case '.':
				row = append(row, 0)
			case '^':
				row = append(row, -1)
			case 'S':
				row = append(row, 1)
			}
		}

		grid = append(grid, row)
	}

	log.Printf("Built grid: %v", grid)

	return grid
}

func (s *Server) Day7Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	grid := buildSplitGrid(req.GetData())

	splits := 0
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// Split
			if grid[i][j] == -1 {
				if grid[i-1][j] == 1 {
					splits++
					grid[i][j-1] = 1
					grid[i][j+1] = 1
				}
			} else if grid[i-1][j] == 1 {
				grid[i][j] = 1
			}
		}
	}

	log.Printf("Got to %v", grid[len(grid)-1])

	return &pb.SolveResponse{Answer: int32(splits)}, nil
}

func (s *Server) Day7Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	grid := buildSplitGrid(req.GetData())

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// Split
			if grid[i][j] == -1 {
				if grid[i-1][j] > 0 {
					grid[i][j-1] += grid[i-1][j]
					grid[i][j+1] += grid[i-1][j]
				}
			} else if grid[i-1][j] > 0 {
				grid[i][j] += grid[i-1][j]
			}
		}
	}

	log.Printf("Got to %v", grid[len(grid)-1])

	sumv := int32(0)
	for _, val := range grid[len(grid)-1] {
		sumv += int32(val)
	}

	return &pb.SolveResponse{Answer: int32(sumv)}, nil
}
