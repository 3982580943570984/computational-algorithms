package main

import "log"

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
	n := len(s.Points) - 1
	s.Polynomials = make([]Coefficient, n)

	// Создание и заполнение матрицы коэффициентов системы линейных уравнений
	matrix := make([][]float64, len(s.Points))
	for i := range matrix {
		matrix[i] = make([]float64, n+2)
	}
	for i := 1; i < n; i++ {
		// выражает часть системы, отвечающую за первую производную на границах интервалов, а именно
		// - разность значений аргумента между соседними точками данных
		matrix[i][i-1] = s.Points[i].X - s.Points[i-1].X
		// отражает влияние текущего интервала на кубический полином, а также на его первую производную в пределах интервала
		matrix[i][i] = 2.0 * (s.Points[i+1].X - s.Points[i-1].X)
		// также выражает разность значений аргумента между соседними точками данных, но уже на следующем интервале
		matrix[i][i+1] = s.Points[i+1].X - s.Points[i].X
		// определяет влияние значений функции на границах интервала, а также изменения скорости изменения значения функции на границах интервала
		// выражение соответствует условию непрерывности первой производной на границе интервалов и условию непрерывности
		// второй производной на границах интервалов, устанавливающих значение второй производной равным нулю на границах интервалов
		matrix[i][n+1] = 6.0 * ((s.Points[i+1].Y-s.Points[i].Y)/(s.Points[i+1].X-s.Points[i].X) - (s.Points[i].Y-s.Points[i-1].Y)/(s.Points[i].X-s.Points[i-1].X))
	}
	matrix[0][0] = 1
	matrix[n][n] = 1

	// Решение системы линейных уравнений для нахождения вторых производных
	secondDerivatives := solveTridiagonalSystem(matrix)

	// Вычисление коэффициентов кубических полиномов сплайна
	// https://mathworld.wolfram.com/CubicSpline.html
	for i := 0; i < n; i++ {
		h := s.Points[i+1].X - s.Points[i].X
		s.Polynomials[i] = Coefficient{
			s.Points[i].Y,
			(s.Points[i+1].Y-s.Points[i].Y)/h - h/6*(2*secondDerivatives[i]+secondDerivatives[i+1]),
			secondDerivatives[i] / 2,
			(secondDerivatives[i+1] - secondDerivatives[i]) / (6 * h),
		}
	}
}

/**
solveTridiagonalSystem решает систему линейных уравнений с трехдиагональной матрицей
https://en.wikipedia.org/wiki/Tridiagonal_matrix_algorithm
*/
func solveTridiagonalSystem(matrix [][]float64) []float64 {
	// Получаем размер матрицы
	n := len(matrix) - 1

	// Прямой ход прогонки
	for i := 1; i <= n; i++ {
		// Вычисляем коэффициент m для текущей строки
		m := matrix[i][i-1] / matrix[i-1][i-1]
		// Изменяем элементы матрицы в соответствии с прогонкой
		matrix[i][i] -= m * matrix[i-1][i]
		matrix[i][n+1] -= m * matrix[i-1][n+1]
	}

	// Обратный ход прогонки
	x := make([]float64, n+1)
	x[n] = matrix[n][n+1] / matrix[n][n]
	for i := n - 1; i >= 0; i-- {
		x[i] = (matrix[i][n+1] - matrix[i][i+1]*x[i+1]) / matrix[i][i]
	}

	// Возвращаем решение системы
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
			dx := x - s.Points[i].X
			a := s.Polynomials[i].a
			b := s.Polynomials[i].b
			c := s.Polynomials[i].c
			d := s.Polynomials[i].d
			return a + dx*(b+dx*(c+dx*d))
		}
	}

	return 0.0
}
