package main

import (
	"fmt"
	"math/rand"
)

func sort(data []int) {
	buildMaxHeap(data)

	fmt.Println(data)

	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]

		heapify(data, 0, i-1)

		fmt.Println(data)
	}
}

func buildMaxHeap(data []int) {
	for i := len(data) / 2; i >= 0; i-- {
		heapify(data, i, len(data)-1)
	}
}

func heapify(data []int, targetIdx, heapIdx int) {
	largestIdx := targetIdx

	leftIdx := 2*targetIdx + 1
	if leftIdx <= heapIdx && data[leftIdx] > data[largestIdx] {
		largestIdx = leftIdx
	}

	rightIdx := 2 * (targetIdx + 1)
	if rightIdx <= heapIdx && data[rightIdx] > data[largestIdx] {
		largestIdx = rightIdx
	}

	if largestIdx != targetIdx {
		data[targetIdx], data[largestIdx] = data[largestIdx], data[targetIdx]

		heapify(data, largestIdx, heapIdx)
	}
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
