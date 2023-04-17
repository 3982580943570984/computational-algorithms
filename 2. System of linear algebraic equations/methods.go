package main

import (
	"errors"
	"math"
)

func SolveGaussian(A [][]float64, b []float64) ([]float64, error) {
	n := len(A)

	// Augment the matrix A with the vector b
	Ab := make([][]float64, n)
	for i := range Ab {
		Ab[i] = make([]float64, n+1)
		copy(Ab[i], A[i])
		Ab[i][n] = b[i]
	}

	// Forward elimination
	for k := 0; k < n-1; k++ {
		// Find the pivot row
		maxIdx := k
		for i := k + 1; i < n; i++ {
			if math.Abs(Ab[i][k]) > math.Abs(Ab[maxIdx][k]) {
				maxIdx = i
			}
		}

		// Swap the pivot row with the current row
		if maxIdx != k {
			Ab[k], Ab[maxIdx] = Ab[maxIdx], Ab[k]
		}

		// Eliminate the current column
		for i := k + 1; i < n; i++ {
			factor := Ab[i][k] / Ab[k][k]
			for j := k + 1; j <= n; j++ {
				Ab[i][j] -= factor * Ab[k][j]
			}
		}
	}

	// Backward substitution
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += Ab[i][j] * x[j]
		}
		x[i] = (Ab[i][n] - sum) / Ab[i][i]
	}

	return x, nil
}

func SolveJacobi(A [][]float64, b []float64, tol float64, maxIter int) ([]float64, error) {
	n := len(A)

	// Check that A is diagonally dominant
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			if j != i {
				sum += math.Abs(A[i][j])
			}
		}
		if math.Abs(A[i][i]) <= sum {
			return nil, errors.New("A is not diagonally dominant")
		}
	}

	// Initialize x to zero
	x := make([]float64, n)

	// Iterate until convergence or maximum number of iterations
	for k := 0; k < maxIter; k++ {
		// Compute the next approximation of x
		xNext := make([]float64, n)
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += A[i][j] * x[j]
				}
			}
			xNext[i] = (b[i] - sum) / A[i][i]
		}

		// Check for convergence
		converged := true
		for i := 0; i < n; i++ {
			if math.Abs(xNext[i]-x[i]) > tol {
				converged = false
				break
			}
		}
		if converged {
			return xNext, nil
		}

		// Update x for the next iteration
		x = xNext
	}

	return nil, errors.New("maximum number of iterations exceeded")
}
