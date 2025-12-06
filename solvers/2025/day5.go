package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func readData(data string) ([][]int, []int) {
	var ranges = make([][]int, 0)
	var numbers = make([]int, 0)
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			break
		}
		elems := strings.Split(line, "-")
		start, err := strconv.ParseInt(elems[0], 10, 32)
		if err != nil {
			log.Fatalf("Cannot parse range start: %v -> %v", line, err)
		}
		end, err := strconv.ParseInt(elems[1], 10, 32)
		if err != nil {
			log.Fatalf("Cannot parse range end: %v -> %v", line, err)
		}
		ranges = append(ranges, []int{int(start), int(end)})
	}

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 || strings.Contains(line, "-") {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatalf("Cannot parse line: %v -> %v", line, err)
		}
		numbers = append(numbers, int(num))
	}

	return ranges, numbers
}

func (s *Server) Day5Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	ranges, numbers := readData(req.GetData())

	sumv := 0
	for _, num := range numbers {
		for _, rangev := range ranges {
			if num >= rangev[0] && num <= rangev[1] {
				log.Printf("Found %v in range %v", num, rangev)
				sumv++
				break
			}
		}
	}

	return &pb.SolveResponse{
		Answer: int32(sumv),
	}, nil
}
