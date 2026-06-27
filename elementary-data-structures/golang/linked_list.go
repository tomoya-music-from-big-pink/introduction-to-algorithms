package main

import "fmt"

type Element struct {
	Data       int
	Next, Prev *Element
}

type LinkedList struct {
	Size       int
	Head, Tail *Element
}

func (ll *LinkedList) add(data int) {
	ll.Size++

	z := &Element{data, nil, nil}

	if ll.Head == nil {
		ll.Head = z
	} else {
		ll.Tail.Next = z
	}
	z.Prev = ll.Tail
	ll.Tail = z
}

func (ll *LinkedList) remove(data int) {
	z := ll.contains(data)

	if z == nil {
		fmt.Printf("%d is not found in the list\n", data)

		return
	}

	ll.Size--

	if z.Next != nil {
		z.Next.Prev = z.Prev
	} else {
		ll.Tail = z.Prev
	}
	if z.Prev != nil {
		z.Prev.Next = z.Next
	} else {
		ll.Head = z.Next
	}
}

func (ll *LinkedList) contains(data int) *Element {
	for p := ll.Head; p != nil; p = p.Next {
		if p.Data == data {
			return p
		}
	}

	return nil
}

func (ll *LinkedList) printList() {
	fmt.Printf("size = %d\n", ll.Size)

	if ll.Head == nil {
		fmt.Println("(empty)")

		return
	}

	fmt.Println("--- from head to tail ---")
	for p := ll.Head; p != nil; p = p.Next {
		fmt.Print(p.Data)
		if p.Next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Print("\n")
		}
	}

	fmt.Println("--- from tail to head ---")
	for p := ll.Tail; p != nil; p = p.Prev {
		fmt.Print(p.Data)
		if p.Prev != nil {
			fmt.Print(" <- ")
		} else {
			fmt.Print("\n")
		}
	}
}

func main() {
	LL := &LinkedList{}

	for {
		fmt.Print("1:add 2:remove 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)
			LL.add(data)
		case 2:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)
			LL.remove(data)
		case 3:
			LL.printList()
		default:
			fmt.Println("invalid operation")
		}
	}
}
