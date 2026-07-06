package main

import (
	"container/heap"
	"fmt"
	"math"
)

func prim(graph *Graph, r *Vertex) {
	vertexHeap := make(vertexHeap, 0)
	inHeap := make(map[string]bool)
	verticies := graph.vertices

	for _, vertex := range verticies {
		vertexHeap = append(vertexHeap, vertex)
		inHeap[vertex.name] = true
	}

	for _, v := range vertexHeap {
		if r.name != v.name {
			v.distance = math.MaxInt64
		} else {
			v.distance = 0
		}
	}

	heap.Init(&vertexHeap)

	adjacency_list := graph.adjacencyList
	for len(vertexHeap) > 0 {
		u := heap.Pop(&vertexHeap).(*Vertex)
		inHeap[u.name] = false

		adjacencies := adjacency_list[u.name]
		for _, adjacency := range adjacencies {
			v := verticies[adjacency.end]
			weight := adjacency.weight
			if inHeap[v.name] && v.distance > weight {
				fmt.Printf("u = %s, v = %s / before = %d, after = %d\n", u.name, v.name, v.distance, weight)

				v.distance = weight
				v.predecessor = u

				heap.Init(&vertexHeap)
			}
		}
	}
}

func main() {
	vertices := make(map[string]*Vertex)
	for _, name := range "abcdefghi" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{start: "a", end: "b", weight: 4})
	edges = append(edges, &Edge{start: "a", end: "h", weight: 8})
	edges = append(edges, &Edge{start: "b", end: "a", weight: 4})
	edges = append(edges, &Edge{start: "b", end: "c", weight: 8})
	edges = append(edges, &Edge{start: "b", end: "h", weight: 11})
	edges = append(edges, &Edge{start: "c", end: "b", weight: 8})
	edges = append(edges, &Edge{start: "c", end: "d", weight: 7})
	edges = append(edges, &Edge{start: "c", end: "f", weight: 4})
	edges = append(edges, &Edge{start: "c", end: "i", weight: 2})
	edges = append(edges, &Edge{start: "d", end: "c", weight: 7})
	edges = append(edges, &Edge{start: "d", end: "e", weight: 9})
	edges = append(edges, &Edge{start: "d", end: "f", weight: 14})
	edges = append(edges, &Edge{start: "e", end: "d", weight: 9})
	edges = append(edges, &Edge{start: "e", end: "f", weight: 10})
	edges = append(edges, &Edge{start: "f", end: "c", weight: 4})
	edges = append(edges, &Edge{start: "f", end: "d", weight: 14})
	edges = append(edges, &Edge{start: "f", end: "e", weight: 10})
	edges = append(edges, &Edge{start: "f", end: "g", weight: 2})
	edges = append(edges, &Edge{start: "g", end: "f", weight: 2})
	edges = append(edges, &Edge{start: "g", end: "h", weight: 1})
	edges = append(edges, &Edge{start: "g", end: "i", weight: 6})
	edges = append(edges, &Edge{start: "h", end: "a", weight: 8})
	edges = append(edges, &Edge{start: "h", end: "b", weight: 11})
	edges = append(edges, &Edge{start: "h", end: "g", weight: 1})
	edges = append(edges, &Edge{start: "h", end: "i", weight: 7})
	edges = append(edges, &Edge{start: "i", end: "c", weight: 2})
	edges = append(edges, &Edge{start: "i", end: "g", weight: 6})
	edges = append(edges, &Edge{start: "i", end: "h", weight: 7})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.initializeGraph()

	fmt.Println("--- graph ---")
	graph.printGraph()

	fmt.Println("--- Prim's algorithm ---")
	prim(graph, vertices["a"])

	fmt.Println("--- result ---")
	total := 0
	for _, u := range vertices {
		fmt.Printf("u = %s", u.name)
		if u.predecessor != nil {
			fmt.Printf(", predecessor = %s", u.predecessor.name)
		}
		fmt.Printf(", weight = %d\n", u.distance)
		total += u.distance
	}
	fmt.Printf("total = %d\n", total)
}
