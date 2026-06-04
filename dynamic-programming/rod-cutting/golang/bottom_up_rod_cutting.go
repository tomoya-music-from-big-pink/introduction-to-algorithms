package main

import (
	"fmt"
	"math"
)

func cutRod(prices []int, length int) ([]int, []int) {
	results := make([]int, length+1)
	slices := make([]int, length+1)

	results[0] = 0

	for j := 1; j < length; j++ {
		result := -math.MaxInt64
		for i := 1; i < j+1; i++ {
			if result < prices[i]+results[j-i] {
				result = prices[i] + results[j-i]
				slices[j] = i
			}
		}

		results[j] = result
	}

	return results, slices
}

func printOptimalStrategy(slices []int, length int) {
	fmt.Print("(")

	for rest := length; rest > 0; rest -= slices[rest] {
		fmt.Print(slices[length])

		if rest-slices[rest] > 0 {
			fmt.Print(" + ")
		}
	}

	fmt.Print(")")
}

func main() {
	prices := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

	fmt.Print("  i |")
	for i := range prices {
		fmt.Printf("%3d|", i)
	}
	fmt.Print("\n")
	fmt.Print("p[i]|")
	for _, p := range prices {
		fmt.Printf("%3d|", p)
	}
	fmt.Print("\n")

	results, slices := cutRod(prices, len(prices))
	for i := range prices {
		fmt.Printf("length = %d / result = %d ", i, results[i])

		if i != 0 {
			printOptimalStrategy(slices, i)
		}

		fmt.Print("\n")
	}
}
