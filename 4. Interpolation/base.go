package main

type Point struct {
	X, Y float64
}

// Определение узлов методом равноотстоящих узлов
func EquidistantNodes(a, b float64, n int) ([]float64, []float64) {
	X, Y := []float64{}, []float64{}

	h := (b - a) / float64(n)
	for i := 0; i <= int(n); i++ {
		X = append(X, a+h*float64(i))
		Y = append(Y, f(X[i]))
	}

	return X, Y
}
