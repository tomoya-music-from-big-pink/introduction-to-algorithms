package main

import (
	"fmt"
	"math/rand"
)

func search(data []int, key int) int {
	for i, d := range data {
		if d == key {
			return i
		}
	}

	return -1
}

func main() {
	data := make([]int, 20)

	for i := range 20 {
		data[i] = rand.Intn(20) + 1
	}

	fmt.Println("--- data ---")
	fmt.Println(data)

	fmt.Print("input key > ")
	var key int
	_, err := fmt.Scanf("%d", &key)
	if err != nil {
		fmt.Println("input error")

		return
	}

	result := search(data, key)
	if result != -1 {
		fmt.Printf("result = %d\n", result)
	} else {
		fmt.Printf("%d is not found.\n", key)
	}
}
