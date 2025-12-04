package main

import (
	"context"
	"log"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func readGrid(data string) [][]int {
	lines := strings.Split(data, "\n")
	grid := make([][]int, len(lines))
	for y := range len(lines) {
		grid[y] = make([]int, len(lines[y]))
		for x, ch := range lines[y] {
			if ch == '.' {
				grid[y][x] = 0
			} else {
				grid[y][x] = 1
			}
		}
	}

	return grid
}

func accessable(grid [][]int, x, y int) int32 {
	sumv := 0
	for dy := y - 1; dy <= y+1; dy++ {
		for dx := x - 1; dx <= x+1; dx++ {
			if dy >= 0 && dy < len(grid) && dx >= 0 && dx < len(grid[dy]) {
				sumv += grid[dy][dx]
			}
		}
	}

	log.Printf("%v,%v -> %v", x, y, sumv)

	if sumv > 4 {
		return 0
	}
	return 1
}

func (s *Server) Day4Part1(_ context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	grid := readGrid(req.GetData())

	sum := int32(0)
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 1 {
				sum += accessable(grid, x, y)
			}
		}
	}

	return &pb.SolveResponse{Answer: sum}, nil
}
