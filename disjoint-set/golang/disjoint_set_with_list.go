package main

import "fmt"

type Element struct {
	data       string
	next, prev *Element
}

type LinkedList struct {
	head, tail *Element
}

func makeSet(data string) *LinkedList {
	ll := &LinkedList{}

	e := &Element{data: data}
	ll.head = e
	ll.tail = e

	return ll
}

func printSets(sets map[string]*LinkedList) {
	for k, v := range sets {
		fmt.Printf("%v : ", k)

		for p := v.head; p != nil; p = p.next {
			fmt.Printf("%v", p.data)
			if p.next != nil {
				fmt.Print(" -> ")
			}
		}

		fmt.Print("\n")
	}
}

func findSet(ch string, sets map[string]*LinkedList) string {
	for _, v := range sets {
		found := false
		for p := v.head; p != nil; p = p.next {
			if p.data == ch {
				found = true

				break
			}
		}

		if found {
			return v.head.data
		}
	}

	return ""
}

func union(ch1, ch2 string, sets map[string]*LinkedList) {
	parentOfX := findSet(ch1, sets)
	parentOfY := findSet(ch2, sets)

	if parentOfX == parentOfY {
		return
	}

	x := sets[parentOfX]
	y := sets[parentOfY]

	if parentOfX < parentOfY {
		x.tail.next = y.head
		x.tail = y.tail
		delete(sets, parentOfY)
	}
}

func main() {
	sets := make(map[string]*LinkedList)

	for _, ch := range "abcdefghi" {
		sets[string(ch)] = makeSet(string(ch))
	}

	fmt.Println("--- initialize ---")
	printSets(sets)

	fmt.Println("union b and d")
	union("b", "d", sets)
	printSets(sets)
	fmt.Println("union e and g")
	union("e", "g", sets)
	printSets(sets)
	fmt.Println("union a and c")
	union("a", "c", sets)
	printSets(sets)
	fmt.Println("union h and i")
	union("h", "i", sets)
	printSets(sets)
	fmt.Println("union a and b")
	union("a", "b", sets)
	printSets(sets)
	fmt.Println("union e and f")
	union("e", "f", sets)
	printSets(sets)
	fmt.Println("union b and c")
	union("b", "c", sets)
	printSets(sets)
}
