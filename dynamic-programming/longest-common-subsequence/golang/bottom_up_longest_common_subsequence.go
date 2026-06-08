package main

import (
	"fmt"
	"math"
)

type Direction int

const (
	UPPER    Direction = 0
	LEFT     Direction = 1
	DIAGONAL Direction = 2
)

func common_lcs(x, y string) ([][]int, [][]Direction) {
	results := make([][]int, len(x))
	directions := make([][]Direction, len(x))

	for i := range len(x) {
		results[i] = make([]int, len(y))
		directions[i] = make([]Direction, len(y))
		for j := range len(y) {
			results[i][j] = -math.MaxInt64
			directions[i][j] = -math.MaxInt64
		}
	}

	for i := range len(x) {
		for j := range len(y) {
			if i == 0 && j == 0 {
				if x[i] == y[j] {
					results[i][j] = 1
					directions[i][j] = DIAGONAL
				} else {
					results[i][j] = 0
					directions[i][j] = UPPER
				}
			} else if i == 0 {
				if x[i] == y[j] {
					results[i][j] = 1
					directions[i][j] = DIAGONAL
				} else {
					results[i][j] = results[i][j-1]
					directions[i][j] = LEFT
				}
			} else if j == 0 {
				if x[i] == y[j] {
					results[i][j] = 1
					directions[i][j] = DIAGONAL
				} else {
					results[i][j] = results[i-1][j]
					directions[i][j] = UPPER
				}
			} else {
				if x[i] == y[j] {
					results[i][j] = results[i-1][j-1] + 1
					directions[i][j] = DIAGONAL
				} else {
					if results[i-1][j] < results[i][j-1] {
						results[i][j] = results[i][j-1]
						directions[i][j] = LEFT
					} else {
						results[i][j] = results[i-1][j]
						directions[i][j] = UPPER
					}
				}
			}
		}
	}

	return results, directions
}

func printOptimalStrategy(x string, directions [][]Direction, i, j int) {
	if i < 0 || j < 0 {
		return
	}

	if directions[i][j] == UPPER {
		printOptimalStrategy(x, directions, i-1, j)
	} else if directions[i][j] == LEFT {
		printOptimalStrategy(x, directions, i, j-1)
	} else {
		printOptimalStrategy(x, directions, i-1, j-1)
		fmt.Printf("%c", x[i])
	}
}

func main() {
	x := "ABCBDAB"
	y := "BDCABA"

	results, directions := common_lcs(x, y)

	for i := 0; i < len(x); i++ {
		fmt.Printf("--- i = %d ---\n", i)

		for j := 0; j < len(y); j++ {
			fmt.Printf("i = %d, j = %d, x = %v, y = %v / result = %d ", i, j, x[:i+1], y[:j+1], results[i][j])
			printOptimalStrategy(x, directions, i, j)
			fmt.Print("\n")
		}
	}
}
