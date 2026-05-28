package main

import (
	"fmt"
	"math/rand"
)

func sort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		idxOfMin := i

		for j := i + 1; j < len(data); j++ {
			if data[j] < data[idxOfMin] {
				idxOfMin = j
			}
		}

		if idxOfMin != i {
			data[i], data[idxOfMin] = data[idxOfMin], data[i]
		}

		fmt.Println(data)
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
