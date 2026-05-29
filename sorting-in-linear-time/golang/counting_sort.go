package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func sort(data []int) {
	countingTable := make([]int, slices.Max(data)+1)

	for _, d := range data {
		countingTable[d] += 1
	}
	for i := 0; i < len(countingTable)-1; i++ {
		countingTable[i+1] += countingTable[i]
	}

	fmt.Println("--- counting table ---")
	fmt.Println(countingTable)

	temp := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		idxToReplace := countingTable[data[i]] - 1
		temp[idxToReplace] = data[i]
		countingTable[data[i]] -= 1
	}

	clear(data)
	data = data[:0]

	data = append(data, temp...)
}

func main() {
	data := make([]int, 40)

	for i := range 40 {
		data[i] = rand.Intn(20) + 1
	}

	fmt.Println("--- before ---")
	fmt.Println(data)

	fmt.Println("--- sort ---")
	sort(data)

	fmt.Println("--- after ---")
	fmt.Println(data)
}
