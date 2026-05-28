package main

import (
	"fmt"
)

func search(data []int, key int) int {
	startIdx := 0
	endIdx := len(data) - 1

	for startIdx <= endIdx {
		middleIdx := (startIdx + endIdx) / 2

		if data[middleIdx] == key {
			return middleIdx
		} else if data[middleIdx] < key {
			startIdx = middleIdx + 1
		} else {
			endIdx = middleIdx - 1
		}
	}

	return -1
}

func main() {
	data := make([]int, 20)

	for i := range 20 {
		data[i] = i * i
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
