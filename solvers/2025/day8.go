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

func buildDistanceGridAndArray(data string) []*dist {
	var numarr [][]int64
	var resarr []*dist

	for _, line := range strings.Split(data, "\n") {
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

	return resarr
}

func (s *Server) Day8Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return runDay8Part1(req, 10)
}

func runDay8Part1(req *pb.SolveRequest, maxv int) (*pb.SolveResponse, error) {
	sumv := int32(0)

	distGrid := buildDistanceGridAndArray(req.GetData())

	sort.SliceStable(distGrid, func(x, y int) bool {
		return distGrid[x].dist < distGrid[y].dist
	})

	var circuits [][]int

	for i := range maxv {
		placed := false
		for _, circuit := range circuits {
			for _, val := range circuit {
				if val == distGrid[i].index1 || val == distGrid[i].index2 {
					placed = true
					circuit = append(circuit, distGrid[i].index1)
					circuit = append(circuit, distGrid[i].index2}
					break
				}
			}
			if placed {
				break
			}
		}

		if !placed {
			circuits = append(circuits, []int{distGrid[i].index1, distGrid[i].index2})
		}
	}

	log.Printf("%v", circuits)

	return &pb.SolveResponse{Answer: sumv}, nil
}
