package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	ch                  byte
	freq                int
	left, right, parent *Element
}

type ElementHeap []*Element

func (p ElementHeap) Len() int {
	return len(p)
}

func (p ElementHeap) Less(i, j int) bool {
	return p[i].freq < p[j].freq
}

func (p ElementHeap) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}

func (p *ElementHeap) Push(q any) {
	*p = append(*p, q.(*Element))
}

func (p *ElementHeap) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]

	return item
}

func huffman(elementHeap *ElementHeap) any {
	heap.Init(elementHeap)

	n := elementHeap.Len()
	for i := 0; i < n-1; i++ {
		x := heap.Pop(elementHeap)
		y := heap.Pop(elementHeap)
		fmt.Printf("x: %c - %d\n", x.(*Element).ch, x.(*Element).freq)
		fmt.Printf("y: %c - %d\n", y.(*Element).ch, y.(*Element).freq)

		z := &Element{freq: x.(*Element).freq + y.(*Element).freq}
		z.left = x.(*Element)
		z.right = y.(*Element)
		x.(*Element).parent = z
		y.(*Element).parent = z

		heap.Push(elementHeap, z)
	}

	return heap.Pop(elementHeap)
}

func printTree(root *Element) {
	printTreeInternal(root, 0)
}

func printTreeInternal(r *Element, spaceSize int) {
	if r == nil {
		return
	}

	printTreeInternal(r.left, spaceSize+1)
	for range spaceSize {
		fmt.Print("  ")
	}
	fmt.Printf("%c(%d)\n", r.ch, r.freq)
	printTreeInternal(r.right, spaceSize+1)
}

func main() {
	characters := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	frequencies := []int{45, 13, 12, 16, 9, 5}

	fmt.Print(" ch |")
	for _, ch := range characters {
		fmt.Printf("%4c|", ch)
	}
	fmt.Print("\n")
	fmt.Print("freq|")
	for _, freq := range frequencies {
		fmt.Printf("%4d|", freq)
	}
	fmt.Print("\n")

	elements := make([]Element, len(characters))
	elementHeap := &ElementHeap{}
	for i := 0; i < len(characters); i++ {
		elements[i] = Element{ch: characters[i], freq: frequencies[i]}
		elementHeap.Push(&elements[i])
	}

	root := huffman(elementHeap).(*Element)
	printTree(root)
}
