package main

import "fmt"

type Element struct {
	Data int
	Next *Element
}

type Queue struct {
	Size       int
	Head, Tail *Element
}

func (q *Queue) enqueue(data int) {
	q.Size++

	z := &Element{data, nil}

	if q.Head == nil {
		q.Head = z
	} else {
		q.Tail.Next = z
	}
	q.Tail = z
}

func (Q *Queue) dequeue() (int, bool) {
	if Q.Head == nil {
		return 0, false
	}

	Q.Size--

	dequeued := Q.Head
	Q.Head = dequeued.Next

	return dequeued.Data, true
}

func (Q *Queue) printQueue() {
	fmt.Printf("size = %d : ", Q.Size)

	if Q.Head == nil {
		fmt.Println("(empty)")

		return
	}

	for p := Q.Head; p != nil; p = p.Next {
		fmt.Printf("%d", p.Data)
		if p.Next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Print("\n")
		}
	}
}

func main() {
	Q := &Queue{}

	for {
		fmt.Print("1:enqueue 2:dequeue 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input data > ")
			var data int
			fmt.Scanf("%d", &data)
			Q.enqueue(data)
		case 2:
			dequeued, empty := Q.dequeue()
			if !empty {
				fmt.Println("queue is empty")
			} else {
				fmt.Printf("dequeued data : %d\n", dequeued)
			}
		case 3:
			Q.printQueue()
		default:
			fmt.Println("invalid operation")
		}
	}
}
