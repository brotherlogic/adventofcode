package main

import (
	"context"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

type searchCache struct {
	cache map[string]int64
}

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

func (c *searchCache) runSearch(point string, seen []string, mapping map[string][]string, req []string, goal string) int64 {

	if point == goal {
		for _, r := range req {
			found := false
			for _, seen := range seen {
				if r == seen {
					found = true
					break
				}
			}
			if !found {
				return 0
			}
		}
		return 1
	}

	if _, ok := c.cache[point]; ok {
		//log.Printf("FROM CACHE: %v -> %v", point, c.cache[point])
		return c.cache[point]
	}

	// Don't loop
	for _, s := range seen {
		if point == s {
			return 0
		}
	}

	sumv := int64(0)
	for _, dest := range mapping[point] {
		sumv += c.runSearch(dest, copyArr(seen, point), mapping, req, goal)
	}

	//log.Printf("TO CACHE: %v -> %v", point, sumv)
	c.cache[point] = sumv
	return sumv
}

func buildReverseMapping(mapping map[string][]string) map[string][]string {
	revMapping := make(map[string][]string)
	for k, v := range mapping {
		for _, dest := range v {
			revMapping[dest] = append(revMapping[dest], k)
		}
	}
	return revMapping
}

func (*Server) Day11Part1(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	count := int64(0)

	mapping := buildMapping(req.GetData())
	revMapping := buildReverseMapping(mapping)

	cache := &searchCache{
		cache: make(map[string]int64),
	}
	count = cache.runSearch("out", []string{}, revMapping, []string{}, "you")

	return &pb.SolveResponse{BigAnswer: count}, nil
}

func (*Server) Day11Part2(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	count := int64(0)

	mapping := buildMapping(req.GetData())
	revMapping := buildReverseMapping(mapping)

	cache := &searchCache{
		cache: make(map[string]int64),
	}
	c1 := cache.runSearch("dac", []string{}, revMapping, []string{}, "fft")

	cache = &searchCache{
		cache: make(map[string]int64),
	}
	c2 := cache.runSearch("fft", []string{}, revMapping, []string{}, "svr")

	cache = &searchCache{
		cache: make(map[string]int64),
	}
	c3 := cache.runSearch("out", []string{}, revMapping, []string{}, "dac")

	count = c1 * c2 * c3

	return &pb.SolveResponse{BigAnswer: count}, nil
}
