package main

import "fmt"

type Element struct {
	Data int
	Next *Element
}

type Stack struct {
	Size int
	Head *Element
}

func (S *Stack) push(e *Element) {
	S.Size++

	e.Next = S.Head
	S.Head = e
}

func (S *Stack) pop() (int, bool) {
	if S.Head == nil {
		return 0, false
	}

	S.Size--

	popped := S.Head
	result := popped.Data
	S.Head = popped.Next

	return result, true
}

func (S *Stack) printStack() {
	fmt.Printf("size = %d : ", S.Size)

	if S.Head == nil {
		fmt.Println("(empty)")

		return
	}

	for p := S.Head; p != nil; p = p.Next {
		fmt.Print(p.Data)
		if p.Next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Print("\n")
		}
	}
}

func main() {
	S := &Stack{}

	for {
		fmt.Print("1:push 2:pop 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)
			S.push(&Element{data, nil})
		case 2:
			popped, empty := S.pop()
			if !empty {
				fmt.Println("stack is empty")
			} else {
				fmt.Printf("popped data = %d\n", popped)
			}
		case 3:
			S.printStack()
		default:
			fmt.Println("invalid operation")
		}
	}
}
