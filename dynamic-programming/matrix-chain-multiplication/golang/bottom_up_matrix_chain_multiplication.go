package main

import (
	"fmt"
	"math"
)

func matrixChanOrder(p []int) ([][]int, [][]int) {
	results := make([][]int, len(p))
	thresholds := make([][]int, len(p))

	for i := 0; i < len(p); i++ {
		results[i] = make([]int, len(p))
		thresholds[i] = make([]int, len(p))
		for j := 0; j < len(p); j++ {
			if i == j {
				results[i][j] = 0
				thresholds[i][j] = i
			} else {
				results[i][j] = math.MaxInt64
			}
		}
	}

	for diff := 1; diff < len(p)-1; diff++ {
		for i, j := 0, diff; j < len(p)-1; j++ {
			result := math.MaxInt64
			for k := i; k < j; k++ {
				if result > results[i][k]+results[k+1][j]+p[i]*p[k+1]*p[j+1] {
					result = results[i][k] + results[k+1][j] + p[i]*p[k+1]*p[j+1]
					thresholds[i][j] = k
				}
			}

			results[i][j] = result

			i++
		}
	}

	return results, thresholds
}

func printOptimalStrategy(thresholds [][]int, i, j int) {
	if i == j {
		fmt.Printf("A%d", thresholds[i][j])
	} else {
		fmt.Print("(")
		printOptimalStrategy(thresholds, i, thresholds[i][j])
		printOptimalStrategy(thresholds, thresholds[i][j]+1, j)
		fmt.Print(")")
	}
}

func main() {
	p := []int{30, 35, 15, 5, 10, 20, 25}

	fmt.Print("  i |")
	for i := 0; i < len(p)-1; i++ {
		fmt.Printf("%7d|", i)
	}
	fmt.Print("\n")
	fmt.Print("A[i]|")
	for i := 0; i < len(p)-1; i++ {
		fmt.Printf("%3d*%3d|", p[i], p[i+1])
	}
	fmt.Print("\n")

	results, thresholds := matrixChanOrder(p)

	for diff := 0; diff < len(p)-1; diff++ {
		fmt.Printf("--- diff = %d ---\n", diff)

		for i, j := 0, diff; j < len(p)-1; j++ {
			fmt.Printf("i = %d, j = %d / result = %d ", i, j, results[i][j])

			if i != j {
				printOptimalStrategy(thresholds, i, j)
			}

			fmt.Print("\n")

			i++
		}
	}
}
