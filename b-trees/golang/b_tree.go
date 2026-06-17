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

func (t *Tree) contains(key int) bool {
	return t.root.containsInternal(key)
}

func (x *Element) containsInternal(key int) bool {
	var i int
	for i = 0; i < x.n && x.keys[i] < key; i++ {
	}

	if i < x.n && x.keys[i] == key {
		return true
	} else {
		if x.isLeaf {
			return false
		} else {
			return x.children[i].containsInternal(key)
		}
	}
}

func (t *Tree) remove(key int) {
	if !t.contains(key) {
		fmt.Printf("%d is not found in the tree.\n", key)

		return
	}

	t.root.removeIntenal(t, key)
}

func (x *Element) removeIntenal(t *Tree, key int) {
	var i int
	for i = 0; i < x.n && x.keys[i] < key; i++ {
	}

	if i < x.n && x.keys[i] == key {
		if x.isLeaf {
			x.keys = append(x.keys[:i], x.keys[i+1:]...)
			x.n--
		} else {
			if x.children[i].n >= t.minDegree {
				y := x.children[i]
				keyToLift := y.keys[len(y.keys)-1]

				y.keys = append([]int{}, append(y.keys[:t.minDegree-1], []int{x.keys[i]}...)...)

				x.keys[i] = keyToLift

				x.children[i].removeIntenal(t, key)
			} else if x.children[i+1].n >= t.minDegree {
				y := x.children[i+1]
				keyToLift := y.keys[0]

				y.keys = append([]int{x.keys[i]}, y.keys[1:]...)

				x.keys[i] = keyToLift

				x.children[i+1].removeIntenal(t, key)
			} else {
				y := x.children[i]
				z := x.children[i+1]

				y.keys = append([]int{}, append(y.keys, []int{x.keys[i]}...)...)
				y.keys = append([]int{}, append(y.keys, z.keys...)...)
				y.n = len(y.keys)

				if !y.isLeaf {
					y.children = append([]*Element{}, append(y.children, z.children...)...)
				}

				x.keys = append(x.keys[:i], x.keys[i+1:]...)
				x.children = append(x.children[:i+1], x.children[i+2:]...)
				x.n = len(x.keys)
				if x.n == 0 {
					t.root = y
				}

				y.removeIntenal(t, key)
			}
		}
	} else {
		if x.children[i].n == t.minDegree-1 {
			if i > 0 && x.children[i-1].n >= t.minDegree {
				y := x.children[i]
				z := x.children[i-1]
				keyToLift := z.keys[len(z.keys)-1]
				keyToDrop := x.keys[i-1]

				y.keys = append([]int{keyToDrop}, y.keys...)
				y.n++
				z.keys = append([]int{}, z.keys[:len(z.keys)-1]...)
				z.n--
				x.keys[i-1] = keyToLift

				if !y.isLeaf {
					childrenToMove := z.children[len(z.children)-1]
					y.children = append([]*Element{childrenToMove}, y.children...)
					z.children = append([]*Element{}, z.children[1:]...)
				}

				y.removeIntenal(t, key)
			} else if i < x.n && x.children[i+1].n >= t.minDegree {
				y := x.children[i]
				z := x.children[i+1]
				keyToLift := z.keys[0]
				keyToDrop := x.keys[i]

				y.keys = append([]int{}, append(y.keys, []int{keyToDrop}...)...)
				y.n++
				z.keys = append([]int{}, z.keys[1:]...)
				z.n--
				x.keys[i] = keyToLift

				if !y.isLeaf {
					childrenToMove := z.children[0]
					y.children = append(y.children, childrenToMove)
					z.children = append([]*Element{}, z.children[1:]...)
				}

				y.removeIntenal(t, key)
			} else {
				if i > 0 {
					y := x.children[i]
					z := x.children[i-1]
					keyToDrop := x.keys[i-1]

					z.keys = append(z.keys, []int{keyToDrop}...)
					z.n++
					z.keys = append(z.keys, append([]int{}, y.keys...)...)
					z.n += y.n

					if !y.isLeaf {
						z.children = append(z.children, append([]*Element{}, y.children...)...)
					}

					x.keys = append(x.keys[:i-1], append([]int{}, x.keys[i:]...)...)
					x.n--
					x.children = append(x.children[:i], append([]*Element{}, x.children[i+1:]...)...)

					if x.n == 0 {
						t.root = z
					}

					z.removeIntenal(t, key)
				} else if i < x.n {
					y := x.children[i]
					z := x.children[i+1]
					keyToDrop := x.keys[i]

					y.keys = append(y.keys, []int{keyToDrop}...)
					y.n++
					y.keys = append(y.keys, append([]int{}, z.keys...)...)
					y.n += z.n

					if !y.isLeaf {
						y.children = append(y.children, append([]*Element{}, z.children...)...)
					}

					x.keys = append(x.keys[:i], append([]int{}, x.keys[i+1:]...)...)
					x.n -= 1
					x.children = append(x.children[:i+1], append([]*Element{}, x.children[i+2:]...)...)

					if x.n == 0 {
						t.root = y
					}

					y.removeIntenal(t, key)
				}
			}
		} else {
			x.children[i].removeIntenal(t, key)
		}
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
	t.insert(2)
	t.insert(8)
	t.insert(6)
	t.insert(0)
	t.insert(4)

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
			t.remove(key)
		case 3:
			t.printTree()
		default:
			fmt.Println("invalid operation")
		}
	}
}
