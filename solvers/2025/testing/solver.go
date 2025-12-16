package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(5, 5, []float64{
		1, 0, 1, 1, 0,
		0, 0, 0, 1, 1,
		1, 1, 0, 1, 1,
		1, 1, 0, 0, 1,
		1, 0, 1, 0, 1,
	})

	b := mat.NewVecDense(5, []float64{
		7,
		5,
		12,
		7,
		2,
	})

	var x mat.VecDense

	// Solve the system A * x = b.
	// The SolveVec function handles various methods like LU decomposition.
	if err := x.SolveVec(a, b); err != nil {
		// Handle cases where the matrix is singular (no unique solution)
		log.Fatal(err)
	}

	// Print the solution vector x
	fmt.Printf("Solution x:\n%.2f\n", mat.Formatted(&x, mat.Prefix(""), mat.Squeeze()))

}
