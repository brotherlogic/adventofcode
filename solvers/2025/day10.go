package main

import (
	"context"
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

func buildSwitches(switches []string) [][]int {
	var res [][]int
	for _, entry := range switches {
		pieces := strings.Split(entry[1:len(entry)-1], ",")
		var comb []int
		for _, piece := range pieces {
			val, err := strconv.ParseInt(piece, 10, 64)
			if err != nil {
				log.Fatalf("Cannot parse: %v;%v -> %v", switches, piece, err)
			}
			comb = append(comb, int(val))
		}
		res = append(res, comb)
	}
	return res
}

func buildJoltage(piece string) []int {
	var res []int
	for _, c := range strings.Split(piece[1:len(piece)-1], ",") {
		val, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse: %v -> %v", piece, err)
		}
		res = append(res, int(val))
	}
	return res
}

func buildLine(line string) ([]bool, [][]int, []int) {
	elems := strings.Fields(strings.TrimSpace(line))

	return buildLights(elems[0]), buildSwitches(elems[1 : len(elems)-2]), buildJoltage(elems[len(elems)-1])
}

type state struct {
	lstate []bool
	count  int32
}

func copy(val []bool) []bool {
	var nval []bool
	for _, entry := range val {
		nval = append(nval, entry)
	}
	return nval
}

func runBest(goal []bool, q []*state, switches [][]int) *state {
	nb := q[0]
	qr := q[1:]
	log.Printf("Running %+v", nb)
	if nb.count > 4 {
		return nil
	}

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
		log.Printf("%v -> %v", na, switchs)
		for _, sv := range switchs {
			na[sv] = !na[sv]
		}

		qr = append(qr, &state{
			lstate: na,
			count:  nb.count + 1,
		})
	}
	return runBest(goal, qr, switches)
}

func computeLine(line string) int32 {
	lights, switches, _ := buildLine(line)
	log.Printf("LINE %v", lights)
	istate := &state{
		lstate: make([]bool, len(lights)),
		count:  0,
	}

	found := runBest(lights, []*state{istate}, switches)
	return found.count
}

func (s *Server) Day10Part1(_ context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int32(0)

	for _, line := range strings.Split(strings.TrimSpace(req.GetData()), "\n") {
		sumv += computeLine(line)
	}

	return &pb.SolveResponse{Answer: sumv}, nil
}
