package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func metaCollapse(ranges [][]int64) [][]int64 {
	ll := len(ranges)
	nr := collapseRanges(ranges)
	for len(nr) != ll {
		ll = len(nr)
		nr = collapseRanges(nr)
	}
	return nr
}

func collapseRanges(ranges [][]int64) [][]int64 {
	var nranges [][]int64
	for len(ranges) > 0 {
		crange := ranges[0]
		log.Printf("Assessing %v (%v from %v)", crange, 0, ranges)
		var tranges [][]int64
		for j := 1; j < len(ranges); j++ {
			log.Printf("Validating against %v", ranges[j])
			found := false
			// Does the top end lie in the middle of crange?
			if ranges[j][1] <= crange[1] && ranges[j][1] >= crange[0] {
				log.Printf("Top end %v in %v", ranges[j], crange)
				if ranges[j][0] < crange[0] {
					crange[0] = ranges[j][0]
					found = true
				}
			}

			// Does the bottom end lie in the middle of crange
			if ranges[j][0] <= crange[1] && ranges[j][0] >= crange[0] {
				log.Printf("Bottom end %v in %v", ranges[j], crange)
				if ranges[j][1] > crange[1] {
					crange[1] = ranges[j][1]
					found = true
				}
			}

			// Does ranges envelop crange
			if ranges[j][0] <= crange[0] && ranges[j][1] >= crange[1] {
				log.Printf("Envelop %v in %v", ranges[j], crange)
				crange = ranges[j]
				found = true
			}

			if !found {
				tranges = append(tranges, ranges[j])
			}
		}

		nranges = append(nranges, crange)
		ranges = tranges
		log.Printf("New ranges: %v", ranges)
	}

	log.Printf("Collapse to: %v", nranges)
	return nranges
}

func readData(data string) ([][]int64, []int64) {
	var ranges = make([][]int64, 0)
	var numbers = make([]int64, 0)
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			break
		}
		elems := strings.Split(line, "-")
		start, err := strconv.ParseInt(elems[0], 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse range start: %v -> %v", line, err)
		}
		end, err := strconv.ParseInt(elems[1], 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse range end: %v -> %v", line, err)
		}
		ranges = append(ranges, []int64{(start), (end)})
	}

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 || strings.Contains(line, "-") {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse line: %v -> %v", line, err)
		}
		numbers = append(numbers, (num))
	}

	return ranges, numbers
}

func (s *Server) Day5Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	ranges, numbers := readData(req.GetData())

	sumv := 0
	for _, num := range numbers {
		for _, rangev := range ranges {
			if num >= rangev[0] && num <= rangev[1] {
				sumv++
				break
			}
		}
	}

	return &pb.SolveResponse{
		Answer: int32(sumv),
	}, nil
}

func (s *Server) Day5Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	ranges, _ := readData(req.GetData())

	sumv := int64(0)
	for _, rangev := range metaCollapse(ranges) {
		sumv += rangev[1] - rangev[0]
	}

	return &pb.SolveResponse{
		BigAnswer: sumv,
	}, nil
}
