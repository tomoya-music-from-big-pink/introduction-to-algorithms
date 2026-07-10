package main

import "fmt"

type Vertex struct {
	name                                  string
	predecessor                           *Vertex
	distance, start_time, end_time, color int
}

type Edge struct {
	start, end string
	weight     int
}

type Adjacency struct {
	end    string
	weight int
}

type Graph struct {
	vertices      map[string]*Vertex
	edges         []*Edge
	adjacencyList map[string][]*Adjacency
}

type UnionFind struct {
	names  string
	parent map[string]string
	rank   map[string]int
}

func (graph *Graph) initializeGraph() {
	graph.adjacencyList = make(map[string][]*Adjacency)
	for _, v := range graph.vertices {
		graph.adjacencyList[v.name] = make([]*Adjacency, 0)
	}
	for _, edge := range graph.edges {
		graph.adjacencyList[edge.start] = append(graph.adjacencyList[edge.start], &Adjacency{end: edge.end, weight: edge.weight})
	}
}

func (graph *Graph) printGraph() {
	for _, v := range graph.vertices {
		fmt.Print(v.name, " : ")

		adjacencies := graph.adjacencyList[v.name]
		for i, adjacency := range adjacencies {
			fmt.Printf("%v(%d)", adjacency.end, adjacency.weight)
			if i < len(adjacencies)-1 {
				fmt.Print(" -> ")
			}
		}

		fmt.Print("\n")
	}
}

func (set *UnionFind) initializeUnionSet(names string) {
	set.names = names
	set.parent = make(map[string]string)
	set.rank = make(map[string]int)
	for _, name := range names {
		set.parent[string(name)] = string(name)
		set.rank[string(name)] = 0
	}
}

func (set *UnionFind) findSet(x string) string {
	if x != set.parent[x] {
		set.parent[x] = set.findSet(set.parent[x])
	}

	return set.parent[x]
}

func (set *UnionFind) union(x, y string) {
	rootOfX := set.findSet(x)
	rootOfY := set.findSet(y)

	if x != y {
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

func (set *UnionFind) getGroup(root string) []string {
	var group []string
	for _, name := range set.names {
		if set.findSet(string(name)) == root {
			group = append(group, string(name))
		}
	}
	return group
}

func (set *UnionFind) allGroups() map[string][]string {
	groups := make(map[string][]string)
	for _, name := range set.names {
		root := set.findSet(string(name))
		groups[root] = append(groups[root], string(name))
	}
	return groups
}

func (set *UnionFind) printSet() {
	for root, members := range set.allGroups() {
		fmt.Printf("root=%v: %v\n", root, members)
	}
}

type vertexHeap []*Vertex

func (v vertexHeap) Len() int {
	return len(v)
}

func (v vertexHeap) Less(i, j int) bool {
	return v[i].distance < v[j].distance
}

func (v vertexHeap) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v *vertexHeap) Push(x interface{}) {
	*v = append(*v, x.(*Vertex))
}

func (v *vertexHeap) Pop() interface{} {
	old := *v
	n := len(old)
	x := old[n-1]
	*v = old[0 : n-1]
	return x
}
