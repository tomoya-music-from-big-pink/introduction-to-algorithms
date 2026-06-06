package main

import (
	"fmt"
	"math"
)

func matrixChanOrder(p []int, startIdx, endIdx int) int {
	results := make([][]int, len(p))

	for i := 0; i < len(p); i++ {
		results[i] = make([]int, len(p))
		for j := 0; j < len(p); j++ {
			results[i][j] = math.MaxInt64
		}
	}

	return matrixChanOrderInternal(p, startIdx, endIdx, results)
}

func matrixChanOrderInternal(p []int, startIdx, endIdx int, results [][]int) int {
	if results[startIdx][endIdx] < math.MaxInt64 {
		return results[startIdx][endIdx]
	}

	result := math.MaxInt64
	if startIdx == endIdx {
		result = 0
	} else {
		for k := startIdx; k < endIdx; k++ {
			result = min(result, matrixChanOrderInternal(p, startIdx, k, results)+matrixChanOrderInternal(p, k+1, endIdx, results)+p[startIdx]*p[k+1]*p[endIdx+1])
		}
	}

	results[startIdx][endIdx] = result

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
