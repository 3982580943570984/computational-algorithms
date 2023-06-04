package main

import (
	"log"
)

// Коэффициенты кубического полинома
type Coefficient struct {
	a, b, c, d float64
}

// Структура сплайна
type Spline struct {
	// Узлы сплайна
	Points []Point
	// Кубические полиномы
	Polynomials []Coefficient
}

func NewSpline(points []Point) *Spline {
	spline := &Spline{Points: points}
	spline.computeCoefficients()
	return spline
}

// computeCoefficients вычисляет коэффициенты кубических полиномов сплайна
func (s *Spline) computeCoefficients() {
	// Количество промежутков
	n := len(s.Points) - 1

	// Для каждого промежутка создаем полином
	s.Polynomials = make([]Coefficient, n)

	// Создаем трехдиагональную матрицу
	matrix := make([][]float64, len(s.Points))
	for i := range matrix {
		matrix[i] = make([]float64, n+2)
	}

	for i := 1; i < n; i++ {
		h_i0 := s.Points[i].X - s.Points[i-1].X
		h_i1 := s.Points[i+1].X - s.Points[i].X

		f_i0 := s.Points[i-1].Y
		f_i1 := s.Points[i].Y
		f_i2 := s.Points[i+1].Y

		// A_i = h_{i}
		matrix[i][i-1] = h_i0
		// C_i = 2*(h_{i} + h_{i+1})
		matrix[i][i] = 2 * (h_i0 + h_i1)
		// B_i = h_{i+1}
		matrix[i][i+1] = h_i1

		// F_i
		matrix[i][n+1] = 6 * ((f_i2-f_i1)/(h_i1) - (f_i1-f_i0)/(h_i0))
	}

	// C_0, B_0
	h_i0 := s.Points[1].X - s.Points[0].X
	h_i1 := s.Points[2].X - s.Points[1].X
	matrix[0][0] = 2 * (h_i1 + h_i0)
	matrix[0][1] = s.Points[1].X - s.Points[0].X

	// A_n, C_n
	h_i0 = s.Points[len(s.Points)-2].X - s.Points[len(s.Points)-3].X
	h_i1 = s.Points[len(s.Points)-1].X - s.Points[len(s.Points)-2].X
	matrix[n][n-1] = s.Points[len(s.Points)-1].X - s.Points[len(s.Points)-2].X
	matrix[n][n] = 2 * (h_i1 + h_i0)

	// Решение системы линейных уравнений
	c := solveTridiagonalMatrix(matrix)

	// Вычисление коэффициентов кубических полиномов сплайна
	for i := 0; i < n; i++ {
		h_i := s.Points[i+1].X - s.Points[i].X

		f_i0 := s.Points[i].Y
		f_i1 := s.Points[i+1].Y

		s.Polynomials[i] = Coefficient{
			// a
			f_i0,
			// b
			(f_i1-f_i0)/h_i - (c[i+1]+2*c[i])*(h_i/6),
			// c
			c[i],
			// d
			(c[i+1] - c[i]) / h_i,
		}
	}
}

/*
solveTridiagonalSystem решает систему линейных уравнений с трехдиагональной матрицей
https://en.wikipedia.org/wiki/Tridiagonal_matrix_algorithm
*/
func solveTridiagonalMatrix(matrix [][]float64) []float64 {
	n := len(matrix)

	// Прямой ход
	for i := 1; i < n; i++ {
		m := matrix[i][i-1] / matrix[i-1][i-1]
		matrix[i][i] -= m * matrix[i-1][i]
		matrix[i][n] -= m * matrix[i-1][n-1]
	}

	// Обратный ход
	x := make([]float64, n)
	x[n-1] = matrix[n-1][n] / matrix[n-1][n-1]
	for i := n - 2; i >= 0; i-- {
		x[i] = (matrix[i][n] - matrix[i][i+1]*x[i+1]) / matrix[i][i]
	}

	return x
}

// Определение значения кубического сплайна в точке x
func (s *Spline) Evaluate(x float64) float64 {
	if x < s.Points[0].X || x > s.Points[len(s.Points)-1].X {
		log.Fatalf("X находится вне области определения сплайна")
	}

	// Определение интервала, в котором находится x
	for i := 0; i < len(s.Points)-1; i++ {
		if x >= s.Points[i].X && x <= s.Points[i+1].X {
			// Вычисление значения кубического полинома в точке x
			h := x - s.Points[i].X

			a := s.Polynomials[i].a
			b := s.Polynomials[i].b
			c := s.Polynomials[i].c
			d := s.Polynomials[i].d

			return a + h*(b+h*(c/2+h*d/6))
		}
	}

	return 0.0
}
