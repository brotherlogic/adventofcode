package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func runZ3(switches [][]int64, goal []int64, aim int64) (int64, bool) {
	log.Printf("RUN %v %v %v", switches, goal, aim)
	f, err := os.CreateTemp("", "z3")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())
	for i := range switches {
		f.WriteString(fmt.Sprintf("(declare-const s%v Int)\n", i+1))
	}

	for i := range goal {
		f.WriteString(fmt.Sprintf("(assert (= %v (+", goal[i]))
		for j, sws := range switches {
			found := false
			for _, val := range sws {
				if val == int64(i) {
					found = true
					break
				}
			}
			if found {
				f.WriteString(fmt.Sprintf(" s%v", j+1))
			}
		}
		f.WriteString(")))\n")
	}

	for j := range switches {
		f.WriteString(fmt.Sprintf("(assert (<= 0 s%v))\n", j+1))
	}

	if aim > 0 {
		f.WriteString(fmt.Sprintf("(assert (= %v (+ ", aim))
		for j := range switches {
			f.WriteString(fmt.Sprintf(" s%v", j+1))
		}
		f.WriteString(")))\n")
	}

	f.WriteString("(check-sat)\n")
	f.WriteString("(get-model)\n")
	f.Close()

	//log.Printf("HERE: %+v", f)
	out, err := exec.Command("z3", f.Name()).CombinedOutput()
	if err != nil {
		log.Printf("Failed to run z3 (%v): %v -> %v", f.Name(), err, string(out))
		return 0, false
	}
	log.Printf("GOT: %v (%v)", string(out), f.Name())
	result := make([]int64, len(switches))
	num := int64(0)
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "define-fun") {
			fields := strings.Fields(line)
			nump, err := strconv.ParseInt(fields[1][1:], 10, 64)
			if err != nil {
				log.Fatalf("Failed to parse: %v", err)
			}
			num = nump

			//log.Printf("FOUND %v", num)
		}
		numl := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(line, ")", ""), " ", ""), "(", "")
		numv, err := strconv.ParseInt(numl, 10, 64)
		if err == nil {
			//log.Printf("GOT %v", numv)
			result[num-1] = numv
		}
	}

	val := int64(0)
	for _, entry := range result {
		if entry < 0 {
			return 0, false
		}
		val += entry
	}
	log.Printf("RET %v (%v) from %v", val, aim, f.Name())
	return val, true
}

func runBestJoltage(goal []int64, switches [][]int64) int64 {
	best, ok := runZ3(switches, goal, -1)
	for ok {
		nbest, nok := runZ3(switches, goal, best-1)
		if nok {
			best = nbest
		}
		ok = nok
	}
	return best
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

func computeJoltage(line string) int64 {
	_, switches, joltage := buildLine(line)
	found := runBestJoltage(joltage, switches)
	return found
}

func (s *Server) Day10Part1(_ context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int32(0)

	for _, line := range strings.Split(strings.TrimSpace(req.GetData()), "\n") {
		sumv += computeLine(line)
	}

	return &pb.SolveResponse{Answer: sumv}, nil
}

func (s *Server) Day10Part2(_ context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	sumv := int64(0)

	for _, line := range strings.Split(strings.TrimSpace(req.GetData()), "\n") {
		jolt := computeJoltage(line)
		log.Printf("SOL %v", jolt)
		sumv += jolt
	}

	return &pb.SolveResponse{BigAnswer: sumv}, nil
}
