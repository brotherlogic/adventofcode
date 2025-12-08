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
				row = append(row, 1)
			case 'S':
				row = append(row, 2)
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
			if grid[i][j] == 1 {
				if grid[i-1][j] == 2 {
					splits++
					grid[i][j-1] = 2
					grid[i][j+1] = 2
				}
			} else if grid[i-1][j] == 2 {
				grid[i][j] = 2
			}
		}
	}

	log.Printf("Got to %v", grid[len(grid)-1])

	return &pb.SolveResponse{Answer: int32(splits)}, nil
}
