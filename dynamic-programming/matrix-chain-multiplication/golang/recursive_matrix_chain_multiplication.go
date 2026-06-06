package main

import (
	"fmt"
	"math"
)

func matrixChanOrder(p []int, startIdx, endIdx int) int {
	if startIdx == endIdx {
		return 0
	}

	result := math.MaxInt64
	for k := startIdx; k < endIdx; k++ {
		result = min(result, matrixChanOrder(p, startIdx, k)+matrixChanOrder(p, k+1, endIdx)+p[startIdx]*p[k+1]*p[endIdx+1])
	}

	return result
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

	for diff := 0; diff < len(p)-1; diff++ {
		fmt.Printf("--- diff = %d ---\n", diff)

		for i, j := 0, diff; j < len(p)-1; j++ {
			result := matrixChanOrder(p, i, j)

			fmt.Printf("i = %d, j = %d / result = %d\n", i, j, result)

			i++
		}
	}
}
