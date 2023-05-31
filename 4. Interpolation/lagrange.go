package main

import "fmt"

/*
https://en.wikipedia.org/wiki/Lagrange_polynomial
*/

type Lagrange struct {
	Nodes []Point
}

func NewLagrange(nodes []Point) *Lagrange {
	return &Lagrange{Nodes: nodes}
}

func (lg *Lagrange) Validate(value float64) error {
	for i := 0; i < len(lg.Nodes); i++ {
		for j := 0; j < len(lg.Nodes); j++ {
			if i != j {
				if lg.Nodes[i].X-lg.Nodes[j].X == 0.0 {
					return fmt.Errorf("Два идентичных значения X. Деление на ноль")
				}
			}
		}
	}

	if value < lg.Nodes[0].X {
		return fmt.Errorf("Значение меньше левой границы указанного диапазона")
	}

	if value > lg.Nodes[len(lg.Nodes)-1].X {
		return fmt.Errorf("Значение больше правой границы указанного диапазона")
	}

	return nil
}

func (lg *Lagrange) Interpolate(value float64) float64 {
	var accumulation float64

	for i := 0; i < len(lg.Nodes); i++ {
		production := lg.Nodes[i].Y
		for j := 0; j < len(lg.Nodes); j++ {
			if j != i {
				production *= (value - lg.Nodes[j].X) / (lg.Nodes[i].X - lg.Nodes[j].X)
			}
		}
		accumulation += production
	}

	return accumulation
}
