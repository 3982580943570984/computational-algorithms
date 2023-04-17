package main

import (
	"fmt"
	"time"
)

func test() {
	n := 200
	matrix := GenerateDiagonallyDominantMatrix(n)
	vector := GenerateVector(n)
	eps := 1e-9

	start := time.Now()
	x1, err := SolveJacobi(matrix, vector, eps, 100000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Solution: ", x1)
	fmt.Println("Time: ", time.Since(start).String())

	start = time.Now()
	x2, _ := SolveGaussian(matrix, vector)
	fmt.Println("Solution:", x2)
	fmt.Println("Time: ", time.Since(start).String())

	fmt.Println(compareArrays(x1, x2))
}
