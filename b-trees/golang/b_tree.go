package main

import "fmt"

type Element struct {
	keys     []int
	n        int
	isLeaf   bool
	children []*Element
}

type Tree struct {
	root      *Element
	minDegree int
}

func (t *Tree) insert(key int) {
	r := t.root
	if r.n == 2*t.minDegree-1 {
		s := &Element{}
		s.children = append(s.children, r)
		t.root = s

		t.splitChild(s, 0)
		t.insertNonFull(s, key)
	} else {
		t.insertNonFull(r, key)
	}
}

func (t *Tree) insertNonFull(x *Element, key int) {
	if x.isLeaf {
		var i int
		for i = x.n - 1; i >= 0 && x.keys[i] > key; i-- {
		}

		i++

		x.keys = append(x.keys[:i], append([]int{key}, x.keys[i:]...)...)
		x.n = len(x.keys)
	} else {
		var i int
		for i = x.n - 1; i >= 0 && x.keys[i] > key; i-- {
		}

		i++

		if x.children[i].n == 2*t.minDegree-1 {
			t.splitChild(x, i)
			if x.keys[i] < key {
				i++
			}
		}

		t.insertNonFull(x.children[i], key)
	}
}

func (t *Tree) splitChild(x *Element, i int) {
	y := x.children[i]
	z := &Element{}

	keyToList := y.keys[t.minDegree-1]

	z.keys = append([]int{}, y.keys[t.minDegree:]...)
	z.n = len(z.keys)
	y.keys = y.keys[:t.minDegree-1]
	y.n = len(y.keys)

	z.isLeaf = y.isLeaf
	if !y.isLeaf {
		z.children = append([]*Element{}, y.children[t.minDegree:]...)
		y.children = y.children[:t.minDegree]
	}

	x.keys = append(x.keys[:i], append([]int{keyToList}, x.keys[i:]...)...)
	x.n = len(x.keys)
	x.children = append(x.children[:i+1], append([]*Element{z}, x.children[i+1:]...)...)
}

func (t *Tree) printTree() {
	t.root.printTreeInternal(0)
}

func (r *Element) printTreeInternal(spaceSize int) {
	if r == nil {
		return
	}

	for range spaceSize {
		fmt.Print("  ")
	}
	fmt.Printf("%v\n", r.keys)

	for i := range len(r.children) {
		r.children[i].printTreeInternal(spaceSize + 1)
	}
}

func main() {
	t := &Tree{root: &Element{isLeaf: true}, minDegree: 2}

	t.insert(5)
	t.insert(9)
	t.insert(3)
	t.insert(7)
	t.insert(1)
	//t.insert(5)

	for {
		fmt.Print("1:insert 2:remove 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)
			t.insert(key)
		case 2:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)
		case 3:
			t.printTree()
		default:
			fmt.Println("invalid operation")
		}
	}
}
