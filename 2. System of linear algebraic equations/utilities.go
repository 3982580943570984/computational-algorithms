package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"time"

	"github.com/olekukonko/tablewriter"
)

func GenerateDiagonallyDominantMatrix(n int) [][]float64 {
	rand.Seed(time.Now().UnixNano())

	// Initialize the matrix with random values [-1, 1]
	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
		for j := range A[i] {
			A[i][j] = rand.Float64()*10 - 1
		}
	}

	// Make the matrix diagonally dominant
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			if j != i {
				sum += math.Abs(A[i][j])
			}
		}
		A[i][i] += sum + 1.0
	}

	return A
}

func GenerateDiagonallyDominantMatrixWithInterval(n int, left, right float64) [][]float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize the matrix with random values [left, right]
	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
		for j := range A[i] {
			A[i][j] = r.Float64()*(right-left) + left
		}
	}

	// Make the matrix diagonally dominant
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			if j != i {
				sum += math.Abs(A[i][j])
			}
		}
		A[i][i] += sum + right
	}

	return A
}

func MakeDiagonallyDominantMatrix(A *[][]float64) {
	// Make the matrix diagonally dominant
	for i := 0; i < len(*A); i++ {
		sum := 0.0
		for j := 0; j < len(*A); j++ {
			if j != i {
				sum += math.Abs((*A)[i][j])
			}
		}
		(*A)[i][i] += sum + 1.0
	}
}

func GenerateVector(n int) []float64 {
	rand.Seed(time.Now().UnixNano())

	// Initialize the vector with random values [-1, 1]
	vector := make([]float64, n)
	for i := range vector {
		vector[i] = rand.Float64()*10 - 1
	}

	return vector
}

func GenerateVectorWithInterval(n int, left, right float64) []float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize the vector with random values [left, right]
	vector := make([]float64, n)
	for i := range vector {
		vector[i] = r.Float64()*(right-left) + left
	}

	return vector
}

func measureTime(f interface{}, args ...interface{}) time.Duration {
	start := time.Now()

	// Convert the input function to a reflect.Value
	fn := reflect.ValueOf(f)

	// Convert the input parameters to a slice of reflect.Values
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// Call the input function with the input parameters
	fn.Call(in)

	return time.Since(start)
}

func subtractArrays(arr1 []float64, arr2 []float64) float64 {
	result := 0.0

	for i := range arr1 {
		result += math.Abs(arr1[i] - arr2[i])
	}

	return result
}

func printArray(array []float64) {
	data := make([][]string, 2)

	data[0] = make([]string, len(array))
	for i := 0; i < len(data[0]); i++ {
		data[0][i] = fmt.Sprintf("X[%d]", i)
	}

	for i := 1; i < len(data); i++ {
		data[i] = make([]string, len(array))
		for j := 0; j < len(data[i]); j++ {
			data[i][j] = fmt.Sprintf("%.4f ", array[j])
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(data[0])
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data[1:])
	table.Render()
}

func printMatrix(matrix [][]float64) {
	data := make([][]string, len(matrix)+1)

	data[0] = make([]string, len(matrix))
	for i := 0; i < len(data[0]); i++ {
		data[0][i] = fmt.Sprintf("A[%d]", i)
	}

	for i := 1; i < len(data); i++ {
		data[i] = make([]string, len(matrix))
		for j := 0; j < len(data[i]); j++ {
			data[i][j] = fmt.Sprintf("%.2f ", matrix[i-1][j])
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(data[0])
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data[1:])
	table.Render()
}

func compareArrays(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if fmt.Sprintf("%.9f", math.Abs(a[i]-b[i])) != "0.000000000" {
			return false
		}
	}

	return true
}

func isArrayEmpty(array []float64) bool { return len(array) == 0 }

func clearConsole() {
	fmt.Print("\033[H\033[2J") // ANSI escape sequence for clearing the console
}
