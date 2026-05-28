package main

import (
	"fmt"
	"math/rand"
)

func sort(data []int) {
	for i := len(data) - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if data[j+1] < data[j] {
				data[j], data[j+1] = data[j+1], data[j]

				swapped = true
			}
		}

		if !swapped {
			break
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
