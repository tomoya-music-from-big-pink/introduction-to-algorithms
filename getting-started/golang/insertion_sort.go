package main

import (
	"fmt"
	"math/rand"
)

func sort(data []int) {
	for j := 1; j < len(data); j++ {
		key := data[j]
		i := j - 1

		for i >= 0 && data[i] > key {
			data[i], data[i+1] = data[i+1], data[i]

			i--
		}

		data[i+1] = key

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
