package main

import (
	"fmt"
	"math/rand"
)

func sort(data []int) {
	mergeSort(data, 0, len(data)-1)
}

func mergeSort(data []int, startIdx, endIdx int) {
	if startIdx < endIdx {
		middleIdx := (startIdx + endIdx) / 2

		mergeSort(data, startIdx, middleIdx)
		mergeSort(data, middleIdx+1, endIdx)

		merge(data, startIdx, middleIdx, endIdx)
	}
}

func merge(data []int, startIdx, middleIdx, endIdx int) {
	firstHalf := append([]int{}, data[startIdx:middleIdx+1]...)
	secondHalf := append([]int{}, data[middleIdx+1:endIdx+1]...)

	i, j := 0, 0
	for k := startIdx; k < endIdx+1; k++ {
		if i < len(firstHalf) && j < len(secondHalf) {
			if firstHalf[i] < secondHalf[j] {
				data[k] = firstHalf[i]
				i++
			} else {
				data[k] = secondHalf[j]
				j++
			}
		} else if i < len(firstHalf) {
			data[k] = firstHalf[i]
			i++
		} else {
			data[k] = secondHalf[j]
			j++
		}
	}

	fmt.Println(data)
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
