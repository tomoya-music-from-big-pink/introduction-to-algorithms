package main

import (
	"fmt"
)

func selectActivities(start, finish []int) []int {
	activities := []int{0}

	finished := finish[0]
	for i := 1; i < len(start); i++ {
		if start[i] > finished {
			activities = append(activities, i)

			finished = finish[i]
		}
	}

	return activities
}

func main() {
	start := []int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	finish := []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}

	fmt.Print("  i |")
	for i := range start {
		fmt.Printf("%3d|", i)
	}
	fmt.Print("\n")
	fmt.Print("s[i]|")
	for _, s := range start {
		fmt.Printf("%3d|", s)
	}
	fmt.Print("\n")
	fmt.Print("f[i]|")
	for _, f := range finish {
		fmt.Printf("%3d|", f)
	}
	fmt.Print("\n")

	activities := selectActivities(start, finish)
	fmt.Printf("%v\n", activities)
}
