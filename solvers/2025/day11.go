package main

import (
	"context"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func buildMapping(data string) map[string][]string {

	mapping := make(map[string][]string)

	for _, line := range strings.Split(data, "\n") {
		nline := strings.TrimSpace(line)
		if len(nline) > 0 {
			fields := strings.Fields(nline)

			mapping[fields[0][0:len(fields[0])-1]] = fields[1:]
		}
	}

	return mapping
}

func copyArr(arr []string, add string) []string {
	var narr []string
	for _, elem := range arr {
		narr = append(narr, elem)
	}
	narr = append(narr, add)
	return narr
}

func runSearch(point string, seen []string, mapping map[string][]string) int32 {

	if point == "out" {
		return 1
	}

	// Don't loop
	/*for _, s := range seen {
		if point == s {
			return 0
		}
	}*/

	sumv := int32(0)
	for _, dest := range mapping[point] {
		sumv += runSearch(dest, copyArr(seen, point), mapping)
	}
	return sumv
}

func (*Server) Day11Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	count := int32(0)

	mapping := buildMapping(req.GetData())

	count = runSearch("you", []string{}, mapping)

	return &pb.SolveResponse{Answer: count}, nil
}
