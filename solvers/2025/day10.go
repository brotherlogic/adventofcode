package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func buildLights(lights string) []bool {
	var ret []bool

	for _, c := range lights[1 : len(lights)-1] {
		if c == '.' {
			ret = append(ret, false)
		} else {
			ret = append(ret, true)
		}
	}

	return ret
}

func buildSwitches(switches []string) [][]int64 {
	var res [][]int64
	for _, entry := range switches {
		pieces := strings.Split(entry[1:len(entry)-1], ",")
		var comb []int64
		for _, piece := range pieces {
			val, err := strconv.ParseInt(piece, 10, 64)
			if err != nil {
				log.Fatalf("Cannot parse: %v;%v -> %v", switches, piece, err)
			}
			comb = append(comb, int64(val))
		}
		res = append(res, comb)
	}
	return res
}

func buildJoltage(piece string) []int64 {
	var res []int64
	for _, c := range strings.Split(piece[1:len(piece)-1], ",") {
		val, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse: %v -> %v", piece, err)
		}
		res = append(res, int64(val))
	}
	return res
}

func buildLine(line string) ([]bool, [][]int64, []int64) {
	elems := strings.Fields(strings.TrimSpace(line))

	return buildLights(elems[0]), buildSwitches(elems[1 : len(elems)-1]), buildJoltage(elems[len(elems)-1])
}

type state struct {
	lstate []bool
	jstate []int64
	count  int32
}

func copy(val []bool) []bool {
	var nval []bool
	for _, entry := range val {
		nval = append(nval, entry)
	}
	return nval
}

func copyj(val []int64) []int64 {
	var nval []int64
	for _, entry := range val {
		nval = append(nval, entry)
	}
	return nval
}

func runBest(goal []bool, q []*state, switches [][]int64, seen map[string]bool) *state {

	for len(q) > 0 {
		nb := q[0]
		q = q[1:]

		if _, ok := seen[fmt.Sprintf("%v", nb.lstate)]; ok {
			continue
		}
		seen[fmt.Sprintf("%v", nb.lstate)] = true

		found := true
		for i := range len(goal) {
			if goal[i] != nb.lstate[i] {
				found = false
				break
			}
		}
		if found {
			return nb
		}

		for _, switchs := range switches {
			na := copy(nb.lstate)
			for _, sv := range switchs {
				na[sv] = !na[sv]
			}

			q = append(q, &state{
				lstate: na,
				count:  nb.count + 1,
			})
		}
	}
	return nil
}

func runBestJoltage(goal []int64, q []*state, switches [][]int64, seen map[string]bool) *state {

	for len(q) > 0 {
		nb := q[0]
		q = q[1:]

		if _, ok := seen[fmt.Sprintf("%v", nb.jstate)]; ok {
			continue
		}
		seen[fmt.Sprintf("%v", nb.jstate)] = true

		found := true
		broken := false
		for i := range len(goal) {
			if goal[i] != nb.jstate[i] {
				found = false
				break
			}
			if goal[i] < nb.jstate[i] {
				broken = true
			}
		}
		if found {
			return nb
		}
		if broken {
			continue
		}

		for _, switchs := range switches {
			na := copyj(nb.jstate)
			for _, sv := range switchs {
				na[sv]++
			}

			q = append(q, &state{
				jstate: na,
				count:  nb.count + 1,
			})
		}
	}
	return nil
}

func computeLine(line string) int32 {
	lights, switches, _ := buildLine(line)
	istate := &state{
		lstate: make([]bool, len(lights)),
		count:  0,
	}

	found := runBest(lights, []*state{istate}, switches, make(map[string]bool))
	return found.count
}

func computeJoltage(line string) int32 {
	_, switches, joltage := buildLine(line)
	istate := &state{
		jstate: make([]int64, len(joltage)),
		count:  0,
	}

	found := runBestJoltage(joltage, []*state{istate}, switches, make(map[string]bool))
	return found.count
}

func (s *Server) Day10Part1(_ context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int32(0)

	for _, line := range strings.Split(strings.TrimSpace(req.GetData()), "\n") {
		sumv += computeLine(line)
	}

	return &pb.SolveResponse{Answer: sumv}, nil
}

func (s *Server) Day10Part2(_ context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int32(0)

	for _, line := range strings.Split(strings.TrimSpace(req.GetData()), "\n") {
		sumv += computeJoltage(line)
	}

	return &pb.SolveResponse{Answer: sumv}, nil
}
