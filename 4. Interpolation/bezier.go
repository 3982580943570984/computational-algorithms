package main

import (
	"fmt"
	"log"
)

type Bezier struct {
	Nodes []Point
}

func NewBezier(nodes []Point) *Bezier {
	return &Bezier{Nodes: nodes}
}

// Определяет сегмент сплайна, в котором находится заданная точка
func (bezier *Bezier) FindSegment(value float64) (int, error) {
	// Итерируемся по всем сегментам сплайна
	for i := 0; i < len(bezier.Nodes)-1; i++ {
		// Проверка на входимость значения в сегмент
		if value >= bezier.Nodes[i].X && value <= bezier.Nodes[i+1].X {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Заданное значение находится вне одного сегмента сплайна")
}

// Определяет первую производную сплайна в узлах интерполяции
func (bezier *Bezier) CalculateDerivatives() []float64 {
	derivatives := []float64{0}

	// Итерируемся по всем сегментам сплайна
	for i := 1; i < len(bezier.Nodes)-1; i++ {
		// Расстояние между узлами интерполяции по оси X
		x0 := bezier.Nodes[i].X - bezier.Nodes[i-1].X
		x1 := bezier.Nodes[i+1].X - bezier.Nodes[i].X

		// Расстояние между узлами интерполяции по оси Y
		y0 := bezier.Nodes[i].Y - bezier.Nodes[i-1].Y
		y1 := bezier.Nodes[i+1].Y - bezier.Nodes[i].Y

		// Определение производной сплайна в узле, равном середине сегмента
		derivatives = append(derivatives, 3*((y1/x1)-(y0/x0)))
	}

	return append(derivatives, 0)
}

// Определяет контрольные точки каждого сегмента сплайна
func (bezier *Bezier) CalculateControlPoints() []Point {
	controlPoints := []Point{}

	derivatives := bezier.CalculateDerivatives()
	for i := 0; i < len(bezier.Nodes)-1; i++ {
		// Длина сегмента
		h := bezier.Nodes[i+1].X - bezier.Nodes[i].X

		// Точки начала и конца сегмента
		p0 := bezier.Nodes[i]
		p3 := bezier.Nodes[i+1]

		// Определяем точки в сегменте
		p1 := Point{p0.X + (h / 3), p0.Y + (h/3)*derivatives[i]}
		p2 := Point{p3.X - (h / 3), p3.Y - (h/3)*derivatives[i+1]}

		// Записываем определенные точки
		controlPoints = append(controlPoints, p1)
		controlPoints = append(controlPoints, p2)
	}

	return controlPoints
}

func (bezier *Bezier) Interpolate(value float64) float64 {
	controlPoints := bezier.CalculateControlPoints()
	segment, err := bezier.FindSegment(value)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Границы сегмента
	p0 := bezier.Nodes[segment]
	p3 := bezier.Nodes[segment+1]

	// Контрольные точки сегмента
	p1 := controlPoints[segment*2]
	p2 := controlPoints[segment*2+1]

	// Нормализованная координата
	t := (value - p0.X) / (p3.X - p0.X)

	// Вычисляем значение сплайна в заданной точке
	y := (1-t)*(1-t)*(1-t)*p0.Y + 3*(1-t)*(1-t)*t*p1.Y + 3*(1-t)*t*t*p2.Y + t*t*t*p3.Y

	return y
}
