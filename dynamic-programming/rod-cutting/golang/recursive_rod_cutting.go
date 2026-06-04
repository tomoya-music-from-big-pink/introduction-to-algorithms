package main

import (
	"fmt"
	"math"
)

func cutRod(prices []int, length int) int {
	if length == 0 {
		return 0
	}

	result := -math.MaxInt64
	for i := 1; i < length+1; i++ {
		result = max(result, prices[i]+cutRod(prices, length-i))
	}

	return result
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

	for i := range prices {
		result := cutRod(prices, i)

		fmt.Printf("length = %d / result = %d\n", i, result)
	}
}
