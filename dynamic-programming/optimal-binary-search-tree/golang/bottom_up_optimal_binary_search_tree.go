package main

import (
	"fmt"
	"math"
)

func optimalBst(p, q []float64, i, j int) float64 {
	if j == i-1 {
		return q[i-1]
	}

	w := q[i-1]
	for l := i; l < j+1; l++ {
		w = w + p[l] + q[l]
	}

	result := math.MaxFloat64
	for r := i; r <= j; r++ {
		result = min(result, optimalBst(p, q, i, r-1)+optimalBst(p, q, r+1, j)+w)
	}

	return result
}

func main() {
	p := []float64{0.0, 0.15, 0.10, 0.05, 0.10, 0.20}
	q := []float64{0.05, 0.10, 0.05, 0.05, 0.05, 0.10}

	fmt.Print("  i |")
	for i := range p {
		fmt.Printf("%5d|", i)
	}
	fmt.Print("\n")
	fmt.Print("p[i]|")
	for _, _p := range p {
		fmt.Printf("%5.2f|", _p)
	}
	fmt.Print("\n")
	fmt.Print("q[i]|")
	for _, _q := range q {
		fmt.Printf("%5.2f|", _q)
	}
	fmt.Print("\n")

	for diff := -1; diff < len(p)-1; diff++ {
		fmt.Printf("--- diff = %d ---\n", diff)

		i := 1
		for j := i + diff; j < len(p); j++ {
			result := optimalBst(p, q, i, j)

			fmt.Printf("i = %d, j = %d / result = %.2f\n", i, j, result)

			i++
		}
	}
}
