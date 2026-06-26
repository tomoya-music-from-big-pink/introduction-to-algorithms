package main

import "fmt"

type UnionFind struct {
	parent []int
	rank   []int
}

func (set *UnionFind) initializeSet(n int) {
	set.parent = make([]int, n)
	set.rank = make([]int, n)

	for i := range n {
		set.parent[i] = i
		set.rank[i] = 0
	}
}

func (set *UnionFind) union(x, y int) {
	rootOfX := set.findSet(x)
	rootOfY := set.findSet(y)

	if rootOfX != rootOfY {
		if set.rank[rootOfX] < set.rank[rootOfY] {
			set.parent[rootOfY] = rootOfX
		} else if set.rank[rootOfY] < set.rank[rootOfX] {
			set.parent[rootOfX] = rootOfY
		} else {
			set.parent[rootOfY] = rootOfX
			set.rank[rootOfX]++
		}
	}
}

func (set *UnionFind) findSet(x int) int {
	if set.parent[x] != x {
		set.parent[x] = set.findSet(set.parent[x])
	}

	return set.parent[x]
}

func (set *UnionFind) getGroup(root int) []int {
	var group []int
	for i := range len(set.parent) {
		if set.findSet(i) == root {
			group = append(group, i)
		}
	}
	return group
}

func (set *UnionFind) allGroups() map[int][]int {
	groups := make(map[int][]int)
	for i := range len(set.parent) {
		root := set.findSet(i)
		groups[root] = append(groups[root], i)
	}
	return groups
}

func (set *UnionFind) printSet() {
	for root, members := range set.allGroups() {
		fmt.Printf("root=%d: %v\n", root, members)
	}
}

func main() {
	n := 10
	set := &UnionFind{}
	set.initializeSet(n)

	fmt.Println("--- initialize ---")
	set.printSet()
	fmt.Println("union 1 and 3")
	set.union(1, 3)
	set.printSet()
	fmt.Println("union 4 and 6")
	set.union(4, 6)
	set.printSet()
	fmt.Println("union 0 and 2")
	set.union(0, 2)
	set.printSet()
	fmt.Println("union 7 and 8")
	set.union(7, 8)
	set.printSet()
	fmt.Println("union 0 and 1")
	set.union(0, 1)
	set.printSet()
	fmt.Println("union 4 and 5")
	set.union(4, 5)
	set.printSet()
	fmt.Println("union 1 and 2")
	set.union(1, 2)
	set.printSet()
}
