package main

import "fmt"

func common_lcs(x, y string, i, j int) int {
	if i == 0 && j == 0 {
		if x[i] == y[j] {
			return 1
		} else {
			return 0
		}
	} else if i == 0 {
		if x[i] == y[j] {
			return 1
		} else {
			return common_lcs(x, y, i, j-1)
		}
	} else if j == 0 {
		if x[i] == y[j] {
			return 1
		} else {
			return common_lcs(x, y, i-1, j)
		}
	} else {
		if x[i] == y[j] {
			return common_lcs(x, y, i-1, j-1) + 1
		} else {
			return max(common_lcs(x, y, i-1, j), common_lcs(x, y, i, j-1))
		}
	}
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
