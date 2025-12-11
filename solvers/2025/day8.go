package main

import (
	"context"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

type dist struct {
	index1  int
	index2  int
	coords1 []int64
	coords2 []int64
	dist    float64
}

func computeDist(a, b []int64) float64 {
	return math.Sqrt(float64(((a[0] - b[0]) * (a[0] - b[0])) +
		((a[1] - b[1]) * (a[1] - b[1])) +
		((a[2] - b[2]) * (a[2] - b[2]))))
}

func buildDistanceGridAndArray(data string) ([]*dist, int) {
	var numarr [][]int64
	var resarr []*dist

	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		var row []int64
		for _, piece := range strings.Split(line, ",") {
			num, err := strconv.ParseInt(piece, 10, 64)
			if err != nil {
				log.Fatalf("Cannot parse line: %v -> %v", line, err)
			}
			row = append(row, num)
		}
		numarr = append(numarr, row)
	}

	for i := 0; i < len(numarr); i++ {
		for j := i + 1; j < len(numarr); j++ {
			d := &dist{
				index1:  i,
				index2:  j,
				coords1: numarr[i],
				coords2: numarr[j],
				dist:    computeDist(numarr[i], numarr[j]),
			}

			resarr = append(resarr, d)
		}
	}

	return resarr, len(numarr)
}

func (s *Server) Day8Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return runDay8Part1(req, 1000)
}

func collapse(circuits []map[int]bool) []map[int]bool {
	for i := 0; i < len(circuits); i++ {
		for j := i + 1; j < len(circuits); j++ {
			found := false
			for val, _ := range circuits[j] {
				for vval, _ := range circuits[i] {
					if val == vval {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				for val, _ := range circuits[j] {
					circuits[i][val] = true
					delete(circuits[j], val)
				}
			}
		}
	}

	return circuits
}

func (s *Server) Day8Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	distGrid, maxv := buildDistanceGridAndArray(req.GetData())

	sort.SliceStable(distGrid, func(x, y int) bool {
		return distGrid[x].dist < distGrid[y].dist
	})

	var circuits []map[int]bool

	var fv *dist
	for _, dg := range distGrid {
		if len(circuits) == 1 && len(circuits[0]) == maxv {
			break
		}
		fv = dg

		seen1, seen2 := -1, -1
		for i, sets := range circuits {
			for val := range sets {
				if val == dg.index1 {
					seen1 = i
				}
				if val == dg.index2 {
					seen2 = i
				}
			}
		}

		if seen1 >= 0 && seen2 < 0 {
			circuits[seen1][dg.index2] = true
		}

		if seen2 >= 0 && seen1 < 0 {
			circuits[seen2][dg.index1] = true
		}

		if seen1 >= 0 && seen2 >= 0 {
			nc := make(map[int]bool)
			for val := range circuits[seen1] {
				nc[val] = true
			}
			for val := range circuits[seen2] {
				nc[val] = true
			}
			var ncircuits []map[int]bool
			for i := range circuits {
				if i != seen1 && i != seen2 {
					ncircuits = append(ncircuits, circuits[i])
				}
			}
			ncircuits = append(ncircuits, nc)
			circuits = ncircuits
		}

		if seen1 < 0 && seen2 < 0 {
			temp := make(map[int]bool)
			temp[dg.index1] = true
			temp[dg.index2] = true
			circuits = append(circuits, temp)
		}
	}

	return &pb.SolveResponse{BigAnswer: fv.coords1[0] * fv.coords2[0]}, nil
}

func runDay8Part1(req *pb.SolveRequest, maxv int) (*pb.SolveResponse, error) {
	distGrid, _ := buildDistanceGridAndArray(req.GetData())

	sort.SliceStable(distGrid, func(x, y int) bool {
		return distGrid[x].dist < distGrid[y].dist
	})

	var circuits []map[int]bool

	for i := range maxv {
		placed := false
		for _, circuit := range circuits {
			for val, _ := range circuit {
				if val == distGrid[i].index1 || val == distGrid[i].index2 {
					placed = true
					if val == distGrid[i].index1 || val == distGrid[i].index2 {
						circuit[distGrid[i].index1] = true
						circuit[distGrid[i].index2] = true
					}
					break
				}
			}
		}
		if !placed {
			circuits = append(circuits, make(map[int]bool))
			circuits[len(circuits)-1][distGrid[i].index1] = true
			circuits[len(circuits)-1][distGrid[i].index2] = true
		}
	}

	circuits = collapse(circuits)
	circuits = collapse(circuits)

	var sizes []int
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit))
	}
	sort.Ints(sizes)

	return &pb.SolveResponse{Answer: int32(sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3])}, nil
}
