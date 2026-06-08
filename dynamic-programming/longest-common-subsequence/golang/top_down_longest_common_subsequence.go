package main

import (
	"fmt"
	"math"
)

func common_lcs(x, y string, i, j int) int {
	results := make([][]int, max(len(x), len(y))+1)

	for i := range len(results) {
		results[i] = make([]int, len(results))

		for j := range len(results[i]) {
			results[i][j] = -math.MaxInt64
		}
	}

	return common_lcs_internal(x, y, i, j, results)
}

func common_lcs_internal(x, y string, i, j int, results [][]int) int {
	if results[i][j] > -math.MaxInt64 {
		return results[i][j]
	}

	var result int
	if i == 0 && j == 0 {
		if x[i] == y[j] {
			result = 1
		} else {
			result = 0
		}
	} else if i == 0 {
		if x[i] == y[j] {
			result = 1
		} else {
			result = common_lcs_internal(x, y, i, j-1, results)
		}
	} else if j == 0 {
		if x[i] == y[j] {
			result = 1
		} else {
			result = common_lcs_internal(x, y, i-1, j, results)
		}
	} else {
		if x[i] == y[j] {
			result = common_lcs_internal(x, y, i-1, j-1, results) + 1
		} else {
			result = max(common_lcs_internal(x, y, i-1, j, results), common_lcs_internal(x, y, i, j-1, results))
		}
	}

	results[i][j] = result

	return result
}

func main() {
	x := "ABCBDAB"
	y := "BDCABA"

	for i := 0; i < len(x); i++ {
		fmt.Printf("--- i = %d ---\n", i)

		for j := 0; j < len(y); j++ {
			result := common_lcs(x, y, i, j)
			fmt.Printf("i = %d, j = %d, x = %v, y = %v / result = %d\n", i, j, x[:i+1], y[:j+1], result)
		}
	}
}
