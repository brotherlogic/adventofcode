package server

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	pb "github.com/brotherlogic/adventofcode/proto"
	"google.golang.org/grpc"
)

func TestSolution(t *testing.T) {
	ts, err := RunTestSolver(8080)
	if err != nil {
		t.Errorf("Unable to run solver: %v", err)
	}

	ts.done()
}

type TestSolver struct {
	wg *sync.WaitGroup
	gs *grpc.Server
}

func RunTestSolver(pn int) (*TestSolver, error) {
	solver := &TestSolver{wg: &sync.WaitGroup{}}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", pn))
	if err != nil {
		return nil, err
	}
	solver.gs = grpc.NewServer()
	pb.RegisterSolverServiceServer(solver.gs, solver)
	var runError error
	go func() {
		if err := solver.gs.Serve(lis); err != nil {
			runError = err
		}
	}()

	return solver, runError
}

func (ts *TestSolver) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return &pb.SolveResponse{StringAnswer: "answer"}, nil
}

func (ts *TestSolver) done() {
	ts.gs.GracefulStop()
}
