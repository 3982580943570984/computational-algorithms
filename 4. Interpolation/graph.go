package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func graph() {
	// задаем точки кривой Безье
	points := []plotter.XY{
		{0, 0},
		{1, 2},
		{2, -1},
		{3, 1},
	}

	// создаем график
	p := plot.New()

	// задаем свойства графика
	p.Title.Text = "Кривая Безье"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// добавляем точки на график
	scatter, _ := plotter.NewScatter(plotter.XYs(points))
	scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	p.Add(scatter)

	// создаем кривую Безье
	bezier := plotter.NewFunction(func(x float64) float64 {
		var result float64
		n := len(points) - 1
		for i := 0; i <= n; i++ {
			result += float64(factorial(n)) / float64(factorial(i)*factorial(n-i)) *
				points[i].Y * math.Pow(x, float64(i)) * math.Pow(1-x, float64(n-i))
		}
		return result
	})

	// добавляем кривую на график
	bezier.LineStyle.Width = vg.Points(2)
	bezier.LineStyle.Color = color.RGBA{B: 255, A: 255}
	p.Add(bezier)

	// сохраняем график в файл
	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "bezier.png"); err != nil {
		panic(err)
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
