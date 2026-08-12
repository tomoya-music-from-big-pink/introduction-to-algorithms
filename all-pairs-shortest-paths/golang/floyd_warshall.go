package main

import (
	"fmt"
	"math"
)

func floydWarshall(W [][]int) [][]int {
	n := len(W)
	D := make([][]int, n)
	copy(D, W)

	for k := range n {
		for i := range n {
			for j := range n {
				D[i][j] = min(D[i][j], D[i][k]+D[k][j])
			}
		}

		fmt.Printf("--- k = %d ---\n", k)
		printMatrix(D)
	}

	return D
}

func printMatrix(W [][]int) {
	for i := range W {
		for j := range W[i] {
			if W[i][j] != math.MaxInt8 {
				fmt.Printf("%3d", W[i][j])
			} else {
				fmt.Print("  -")
			}
		}

		fmt.Println("")
	}
}

func main() {
	W := [][]int{
		{0, 3, 8, math.MaxInt8, -4},
		{math.MaxInt8, 0, math.MaxInt8, 1, 7},
		{math.MaxInt8, 4, 0, math.MaxInt8, math.MaxInt8},
		{2, math.MaxInt8, -5, 0, math.MaxInt8},
		{math.MaxInt8, math.MaxInt8, math.MaxInt8, 6, 0},
	}

	fmt.Println("--- graph ---")
	printMatrix(W)

	fmt.Println("--- Flyod-Warshall algorithm ---")
	D := floydWarshall(W)

	fmt.Println("--- result ---")
	printMatrix(D)
}
