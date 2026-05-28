package main

import (
	"fmt"
	"math/rand"
)

func sort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, startIdx, endIdx int) {
	if startIdx < endIdx {
		partitionIdx := partition(data, startIdx, endIdx)

		quickSort(data, startIdx, partitionIdx-1)
		quickSort(data, partitionIdx+1, endIdx)
	}
}

func partition(data []int, startIdx, endIdx int) int {
	pivot := data[endIdx]

	i := startIdx
	for j := startIdx; j < endIdx; j++ {
		if data[j] < pivot {
			data[j], data[i] = data[i], data[j]

			i++
		}
	}

	data[i], data[endIdx] = data[endIdx], data[i]

	fmt.Println(data)

	return i
}

func main() {
	data := make([]int, 20)

	for i := range 20 {
		data[i] = rand.Intn(20) + 1
	}

	fmt.Println("--- before ---")
	fmt.Println(data)

	fmt.Println("--- sort ---")
	sort(data)

	fmt.Println("--- after ---")
	fmt.Println(data)
}
