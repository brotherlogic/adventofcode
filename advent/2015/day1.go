package advent2015

import (
	"context"
	"strings"

	pb "github.com/brotherlogic/adventofcode/proto"
)

func ComputeFloor(str string) int32 {
	left := strings.Count(str, "(")
	right := strings.Count(str, ")")

	return int32(left - right)
}

func Solve2015Day1Part1(ctx context.Context, data string) *pb.SolveResponse {

	return &pb.SolveResponse{Answer: ComputeFloor(data)}

}
