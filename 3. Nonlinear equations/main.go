package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Pair struct {
	first  float64
	second float64
}

func f(x float64) float64 {
	return math.Exp(x) - math.Pow(x, 2) - 3.4
}

func findRoots(a, b, n float64) []Pair {
	result := []Pair{}

	h := (b - a) / n
	for i := 0.0; i < n-1; i++ {
		x1, x2 := a+i*h, a+(i+1)*h
		if f(x1)*f(x2) < 0 {
			result = append(result, Pair{first: x1, second: x2})
		}
	}

	return result
}

func bisection(a, b, e float64) (float64, int64, string) {
	start := time.Now()

	iter := math.Ceil(math.Log2((b - a) / e))

	count := 0
	for i := 0; i < int(iter); i++ {
		mid := (a + b) / 2

		if f(mid) == 0 {
			return mid, int64(count), time.Since(start).String()
		}

		if math.Signbit(f(a)) == math.Signbit(f(mid)) {
			a = mid
		} else {
			b = mid
		}
		count++
	}

	return (a + b) / 2, int64(count), time.Since(start).String()
}

func secant(x0, x1, e float64) (float64, int64, string) {
	start := time.Now()
	count := 0
	for math.Abs(x1-x0) > e {
		x2 := x1 - ((x1-x0)/(f(x1)-f(x0)))*f(x1)
		x0, x1 = x1, x2
		count++
	}

	return x1, int64(count), time.Since(start).String()
}

func main() {
	roots := findRoots(-100000, 100000, 10000)

	// Отображаем найденные интервалы
	fmt.Println("Интервалы, содержащие корни")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Левая граница", "Правая граница"})
	for _, root := range roots {
		table.Append([]string{fmt.Sprintf("%f", root.first), fmt.Sprintf("%f", root.second)})
	}
	table.Render()
	fmt.Println()

	fmt.Println("Метод бисекции")
	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Интервал", "Корень", "Итерации", "Время"})
	for _, root := range roots {
		value, iterations, elapsedTime := bisection(root.first, root.second, 1e-14)
		table.Append([]string{
			fmt.Sprintf("[%f, %f]", root.first, root.second),
			fmt.Sprintf("%f", value),
			fmt.Sprintf("%v", iterations),
			elapsedTime,
		})
	}
	table.Render()
	fmt.Println()

	fmt.Println("Метод секущих")
	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Интервал", "Корень", "Итерации", "Время"})
	for _, root := range roots {
		value, iterations, elapsedTime := secant(root.first, root.second, 1e-14)
		table.Append([]string{
			fmt.Sprintf("[%f, %f]", root.first, root.second),
			fmt.Sprintf("%f", value),
			fmt.Sprintf("%v", iterations),
			elapsedTime,
		})
	}
	table.Render()
	fmt.Println()
}
