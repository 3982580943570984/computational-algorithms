package main

import (
	"fmt"
	"log"
	"math"

	"github.com/guptarohit/asciigraph"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Интерполируемая функция
func f(x float64) float64 {
	return math.Cos(x + math.Pow(math.Cos(x), 3))
}

func OutputGraphInConsole(data [][]float64) {
	graph := asciigraph.PlotMany(data, asciigraph.Precision(1), asciigraph.Height(30), asciigraph.Width(50), asciigraph.SeriesColors(
		asciigraph.Red,
		asciigraph.Yellow,
		asciigraph.Green,
	))

	fmt.Println(graph)
}

func OutputGraphInPNG(data [][]Point) {
	// Create a new plot and set its title
	p := plot.New()
	p.Title.Text = "Three Graphs on One Plot"

	pts1 := make(plotter.XYs, len(data[0]))
	for i := range pts1 {
		pts1[i].X = data[0][i].X
		pts1[i].Y = data[0][i].Y
	}

	pts2 := make(plotter.XYs, len(data[1]))
	for i := range pts2 {
		pts2[i].X = data[1][i].X
		pts2[i].Y = data[1][i].Y
	}

	pts3 := make(plotter.XYs, len(data[2]))
	for i := range pts3 {
		pts3[i].X = data[2][i].X
		pts3[i].Y = data[2][i].Y
	}

	// Add the three data sets to the plot
	err := plotutil.AddLinePoints(p,
		"Default", pts1,
		"Lagrange", pts2,
		"Bezier", pts3)
	if err != nil {
		panic(err)
	}

	// Set the style of the lines and points for each data set
	p.Legend.Top = true
	p.Legend.Left = true

	// Save the plot to a PNG file
	err = p.Save(12*vg.Inch, 12*vg.Inch, "plot.png")
	if err != nil {
		panic(err)
	}
}

func main() {
	NumberOfPoints := 10
	// Задаем точки, через которые должен проходить многочлен
	X, Y := EquidistantNodes(3, 6, NumberOfPoints)
	Points := []Point{}
	for i := 0; i < NumberOfPoints; i++ {
		Points = append(Points, Point{X[i], Y[i]})
	}

	fmt.Println(X)

	var x float64
	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)

	// Решение методом Лагранжа
	lagrange := NewLagrange(Points)
	if err := lagrange.Validate(x); err != nil {
		log.Fatalf("Ошибка валидации при работе метода Лагранжа: %v\n", err)
	}
	fmt.Printf("Интерполяционный многочлен: %v\n", lagrange.Interpolate(x))

	// Решение кубическими сплайнами
	bezier := NewBezier(Points)
	fmt.Printf("Интерполяционный кубический сплайн: %v\n", bezier.Interpolate(x))

	GraphNumbers := 3
	data := make([][]Point, GraphNumbers)

	for i := 0; i < len(data); i++ {
		for _, x := range X {
			if i == 0 {
				data[i] = append(data[i], Point{x, f(x)})
			} else if i == 1 {
				data[i] = append(data[i], Point{x, lagrange.Interpolate(x)})
			} else {
				if x >= 6 {
					x -= 0.4
				}
				data[i] = append(data[i], Point{x, bezier.Interpolate(x)})
			}
		}
	}

	OutputGraphInPNG(data)
}
