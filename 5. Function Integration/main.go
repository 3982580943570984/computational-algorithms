package main

import (
	"fmt"
	"math"
)

const (
	a = 0
	b = 1
	n = 10000
)

func f(x float64) float64 {
	return math.Exp(-x) * math.Cos(math.Pow(x, 2))
}

func IntegrateLeftRectangles(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	integral := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i)*h
		integral += f(x) * h
	}

	return integral
}

func IntegrateRightRectangles(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	integral := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i+1)*h
		integral += f(x) * h
	}

	return integral
}

func IntegrateMidRectangles(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	integral := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i)*h + h/2
		integral += f(x) * h
	}

	return integral
}

func IntegrateTrapezoid(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	integral := (f(a) + f(b)) / 2.0

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		integral += f(x)
	}

	integral *= h
	return integral
}

func IntegrateSimpson(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	integral := f(a) + f(b)

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			integral += 2 * f(x)
		} else {
			integral += 4 * f(x)
		}
	}

	integral *= h / 3.0
	return integral
}

/**

Апроксимация - приближение

**/

func main() {
	fmt.Printf("Значение интеграла:%.30f\n", IntegrateLeftRectangles(a, b, n))
	fmt.Printf("Значение интеграла:%.30f\n", IntegrateRightRectangles(a, b, n))
	fmt.Printf("Значение интеграла:%.30f\n", IntegrateMidRectangles(a, b, n))
	fmt.Printf("Значение интеграла:%.30f\n", IntegrateTrapezoid(a, b, n))
	fmt.Printf("Значение интеграла:%.30f\n", IntegrateSimpson(a, b, n))
}
