package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func sort(data []int) {
	table := make([][]int, slices.Max(data)/10+1)

	for _, d := range data {
		table[d/10] = append(table[d/10], d)
	}

	fmt.Println("--- table before sorting ---")
	for i, t := range table {
		fmt.Println(i, ":", t)
	}

	for i := range table {
		slices.Sort(table[i])
	}

	fmt.Println("--- table before after ---")
	for i, t := range table {
		fmt.Println(i, ":", t)
	}

	clear(data)
	data = data[:0]

	for i := range table {
		data = append(data, table[i]...)
	}
}

func main() {
	data := make([]int, 100)

	for i := range 100 {
		data[i] = rand.Intn(100) + 1
	}

	fmt.Println("--- before ---")
	fmt.Println(data)

	fmt.Println("--- sort ---")
	sort(data)

	fmt.Println("--- after ---")
	fmt.Println(data)
}
