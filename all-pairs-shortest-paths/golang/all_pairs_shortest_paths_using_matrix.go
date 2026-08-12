package main

import (
	"fmt"
	"math"
)

func allPairsShortestPaths(W [][]int) [][]int {
	n := len(W)
	L := make([][]int, n)
	copy(L, W)

	for i := range n {
		fmt.Printf("--- i = %d ---\n", i)

		L = extendShortestPaths(L, W)

		printMatrix(L)
	}

	return L
}

func extendShortestPaths(L, W [][]int) [][]int {
	n := len(W)
	newL := make([][]int, n)
	for i := range n {
		newL[i] = make([]int, n)
	}

	for i := range n {
		for j := range n {
			l := math.MaxInt8
			for k := range n {
				l = min(l, L[i][k]+W[k][j])
			}

			newL[i][j] = l
		}
	}

	return newL
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

	L := allPairsShortestPaths(W)

	fmt.Println("--- result ---")
	printMatrix(L)
}
