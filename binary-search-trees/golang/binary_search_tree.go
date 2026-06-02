package main

import "fmt"

type Element struct {
	key                 int
	left, right, parent *Element
}

type Tree struct {
	size, height int
	root         *Element
}

func (t *Tree) put(key int) {
	t.size++

	z := &Element{key: key}
	var x, y *Element
	x = t.root
	y = nil
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	if y == nil {
		t.root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}

	t.height = t.getHeight(t.root)
}

func (t *Tree) remove(key int) {
	z := t.find(key)
	if z == nil {
		fmt.Printf("%d is not found in the tree\n", key)

		return
	}

	t.size--

	if z.left == nil {
		t.transplant(z, z.right)
	} else if z.right == nil {
		t.transplant(z, z.left)
	} else {
		y := t.successor(z)
		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}

	t.height = t.getHeight(t.root)
}

func (t *Tree) find(key int) *Element {
	for p := t.root; p != nil; {
		if p.key == key {
			return p
		}

		if key < p.key {
			p = p.left
		} else {
			p = p.right
		}
	}

	return nil
}

func (t *Tree) successor(z *Element) *Element {
	if z.right != nil {
		p := z.right
		for ; p.left != nil; p = p.left {
		}

		return p
	} else {
		p := z
		q := z.parent
		for q != nil && p == q.right {
			p = q
			q = q.parent
		}

		return q
	}
}

func (t *Tree) transplant(p, q *Element) {
	if p.parent == nil {
		t.root = q
	} else if p == p.parent.left {
		p.parent.left = q
	} else {
		p.parent.right = q
	}

	if q != nil {
		q.parent = p.parent
	}
}

func (t *Tree) printTree() {
	fmt.Printf("size = %d / height = %d\n", t.size, t.height)

	if t.root != nil {
		printTreeInternal(t.root, 0)
	}
}

func printTreeInternal(r *Element, spaceSize int) {
	if r.left != nil {
		printTreeInternal(r.left, spaceSize+1)
	}

	for range spaceSize {
		fmt.Print("    ")
	}
	fmt.Println(r.key)

	if r.right != nil {
		printTreeInternal(r.right, spaceSize+1)
	}
}

func (t *Tree) getHeight(r *Element) int {
	if r == nil {
		return 0
	}

	return getHeightInternal(t.root)
}

func getHeightInternal(r *Element) int {
	if r == nil {
		return -1
	}

	heightOfLeft := getHeightInternal(r.left)
	heightOfRight := getHeightInternal(r.right)

	return max(heightOfLeft, heightOfRight) + 1
}

func main() {
	t := &Tree{}

	t.put(8)
	t.put(3)
	t.put(10)
	t.put(1)
	t.put(6)
	t.put(14)
	t.put(4)
	t.put(7)
	t.put(13)

	for {
		fmt.Print("1:put 2:remove 3:print > ")
		var op int
		fmt.Scanf("%d", &op)

		switch op {
		case 1:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)
			t.put(key)
		case 2:
			fmt.Print("input key > ")
			var key int
			fmt.Scanf("%d", &key)
			t.remove(key)
		case 3:
			t.printTree()
		default:
			fmt.Println("invalid operation")
		}
	}
}
