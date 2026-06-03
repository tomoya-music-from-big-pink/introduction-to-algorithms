package main

import "fmt"

type Color int

const (
	RED   Color = 0
	BLACK Color = 1
)

func (c Color) String() string {
	switch c {
	case RED:
		return "RED"
	case BLACK:
		return "BLACK"
	default:
		return "Unknown"
	}
}

type Element struct {
	key                 int
	left, right, parent *Element
	color               Color
}

var sentinel *Element

func nilElement() *Element {
	if sentinel == nil {
		sentinel = &Element{key: -1, color: BLACK}
	}

	return sentinel
}

type Tree struct {
	size, height int
	root         *Element
}

func (t *Tree) put(key int) {
	t.size++

	z := &Element{key: key, color: RED, left: nilElement(), right: nilElement(), parent: nilElement()}
	var x, y *Element
	x = t.root
	y = nilElement()
	for x != nilElement() {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y
	z.left = nilElement()
	z.right = nilElement()
	if y == nilElement() {
		t.root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}

	t.fixAfterInsertion(z)

	t.height = t.getHeight(t.root)
}

func (t *Tree) remove(key int) {
	z := t.find(key)
	if z == nilElement() {
		fmt.Printf("%d is not found in the tree\n", key)

		return
	}

	t.size--

	y := z
	yColor := y.color
	var x *Element
	if z.left == nilElement() {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nilElement() {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y := t.successor(z)
		yColor = y.color
		x = y.right
		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		} else {
			x.parent = y
		}

		t.transplant(z, y)
		y.color = z.color
		y.left = z.left
		y.left.parent = y
	}

	if yColor == BLACK {
		t.fixAfterDeletion(x)
	}

	t.height = t.getHeight(t.root)
}

func (t *Tree) find(key int) *Element {
	for p := t.root; p != nilElement(); {
		if p.key == key {
			return p
		}

		if key < p.key {
			p = p.left
		} else {
			p = p.right
		}
	}

	return nilElement()
}

func (t *Tree) successor(z *Element) *Element {
	if z.right != nilElement() {
		p := z.right
		for ; p.left != nilElement(); p = p.left {
		}

		return p
	} else {
		p := z
		q := z.parent
		for q != nilElement() && p == q.right {
			p = q
			q = q.parent
		}

		return q
	}
}

func (t *Tree) transplant(p, q *Element) {
	if p.parent == nilElement() {
		t.root = q
	} else if p == p.parent.left {
		p.parent.left = q
	} else {
		p.parent.right = q
	}

	q.parent = p.parent
}

func (t *Tree) printTree() {
	fmt.Printf("size = %d / height = %d\n", t.size, t.height)

	if t.root != nilElement() {
		printTreeInternal(t.root, 0)
	}
}

func printTreeInternal(r *Element, spaceSize int) {
	if r.left != nilElement() {
		printTreeInternal(r.left, spaceSize+1)
	}

	for range spaceSize {
		fmt.Print("    ")
	}
	fmt.Printf("%d (%s)\n", r.key, r.color)

	if r.right != nilElement() {
		printTreeInternal(r.right, spaceSize+1)
	}
}

func (t *Tree) getHeight(r *Element) int {
	if r == nilElement() {
		return 0
	}

	return getHeightInternal(r)
}

func getHeightInternal(r *Element) int {
	if r == nilElement() {
		return -1
	}

	heightOfLeft := getHeightInternal(r.left)
	heightOfRight := getHeightInternal(r.right)

	return max(heightOfLeft, heightOfRight) + 1
}

func (t *Tree) rotateLeft(x *Element) {
	y := x.right

	x.right = y.left
	if y.left != nilElement() {
		y.left.parent = x
	}

	y.parent = x.parent
	if x.parent == nilElement() {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *Tree) rotateRight(x *Element) {
	y := x.left

	x.left = y.right
	if y.right != nilElement() {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == nilElement() {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}

func (t *Tree) fixAfterInsertion(z *Element) {
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == RED {
				z.parent.color = BLACK
				z.parent.parent.color = RED
				y.color = BLACK
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.rotateLeft(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == RED {
				z.parent.color = BLACK
				z.parent.parent.color = RED
				y.color = BLACK
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rotateRight(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rotateLeft(z.parent.parent)
			}
		}
	}

	t.root.color = BLACK
}

func (t *Tree) fixAfterDeletion(x *Element) {
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED {
				// case 1
				x.parent.color = RED
				w.color = BLACK
				t.rotateLeft(x.parent)
				w = x.parent.right
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				// case 2
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					// case 3
					w.left.color = BLACK
					w.color = RED
					t.rotateRight(w)
					w = x.parent.right
				}

				// case 4
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == RED {
				// case 1
				x.parent.color = RED
				w.color = BLACK
				t.rotateRight(x.parent)
				w = x.parent.left
			}

			if w.right.color == BLACK && w.left.color == BLACK {
				// case 2
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					// case 3
					w.right.color = BLACK
					w.color = RED
					t.rotateLeft(w)
					w = x.parent.left
				}

				// case 4
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}

	x.color = BLACK
}

func main() {
	t := &Tree{root: nilElement()}

	t.put(2)
	t.put(1)
	t.put(4)
	t.put(5)
	t.put(9)
	t.put(3)
	t.put(6)
	t.put(7)

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
