package main

import (
	"fmt"
	"math"
)

const (
	a = 0
	b = 1
	n = 10
	e = 1e-8
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

func RungeRule(I1, I2, O float64) float64 {
	return math.Abs(I2-I1) * (1 / O)
}

/**

Апроксимация - приближение

**/

func main() {
	multiplier := 2
	I1, I2 := IntegrateLeftRectangles(a, b, n), IntegrateLeftRectangles(a, b, multiplier*n)
	for RungeRule(I1, I2, 1) > e {
		multiplier *= 2
		I1, I2 = I2, IntegrateLeftRectangles(a, b, multiplier*n)
	}
	fmt.Printf("Значение интеграла:%.30f\n", I2)

	multiplier = 2
	I1, I2 = IntegrateRightRectangles(a, b, n), IntegrateRightRectangles(a, b, multiplier*n)
	for RungeRule(I1, I2, 1) > e {
		multiplier *= 2
		I1, I2 = I2, IntegrateRightRectangles(a, b, multiplier*n)
	}
	fmt.Printf("Значение интеграла:%.30f\n", I2)

	multiplier = 2
	I1, I2 = IntegrateMidRectangles(a, b, n), IntegrateMidRectangles(a, b, multiplier*n)
	for RungeRule(I1, I2, 3) > e {
		multiplier *= 2
		I1, I2 = I2, IntegrateMidRectangles(a, b, multiplier*n)
	}
	fmt.Printf("Значение интеграла:%.30f\n", I2)

	multiplier = 2
	I1, I2 = IntegrateTrapezoid(a, b, n), IntegrateTrapezoid(a, b, multiplier*n)
	for RungeRule(I1, I2, 3) > e {
		multiplier *= 2
		I1, I2 = I2, IntegrateTrapezoid(a, b, multiplier*n)
	}
	fmt.Printf("Значение интеграла:%.30f\n", I2)

	multiplier = 2
	I1, I2 = IntegrateSimpson(a, b, n), IntegrateSimpson(a, b, multiplier*n)
	for RungeRule(I1, I2, 15) > e {
		multiplier *= 2
		I1, I2 = I2, IntegrateSimpson(a, b, multiplier*n)
	}
	fmt.Printf("Значение интеграла:%.30f\n", I2)
}
