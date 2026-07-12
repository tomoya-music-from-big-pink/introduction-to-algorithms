package main

import (
	"container/heap"
	"fmt"
	"math"
)

func dijkstra(g *Graph, s *Vertex) {
	vertexHeap := make(vertexHeap, 0)
	vertices := g.vertices

	for _, v := range vertices {
		vertexHeap = append(vertexHeap, v)
	}

	for _, v := range vertexHeap {
		if v.name != s.name {
			v.distance = math.MaxInt64
		} else {
			v.distance = 0
		}
	}

	heap.Init(&vertexHeap)

	adjacencyList := g.adjacencyList
	for len(vertexHeap) > 0 {
		u := heap.Pop(&vertexHeap).(*Vertex)
		adjacencies := adjacencyList[u.name]
		for _, adjacency := range adjacencies {
			v := vertices[adjacency.end]
			weight := adjacency.weight
			relax(u, v, weight)

			heap.Init(&vertexHeap)
		}
	}
}

func printReseult(s *Vertex, v *Vertex, v_name string) {
	if v == s {
		fmt.Printf("%s", s.name)
		if v.name != v_name {
			fmt.Print(" -> ")
		}
	} else if v.predecessor != nil {
		printReseult(s, v.predecessor, v_name)
		fmt.Printf("%s", v.name)
		if v.name != v_name {
			fmt.Print(" -> ")
		}
	}
}

func main() {
	vertices := make(map[string]*Vertex)
	for _, name := range "styxz" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{start: "s", end: "t", weight: 10})
	edges = append(edges, &Edge{start: "s", end: "y", weight: 5})
	edges = append(edges, &Edge{start: "t", end: "y", weight: 2})
	edges = append(edges, &Edge{start: "t", end: "x", weight: 1})
	edges = append(edges, &Edge{start: "y", end: "t", weight: 3})
	edges = append(edges, &Edge{start: "y", end: "x", weight: 9})
	edges = append(edges, &Edge{start: "y", end: "z", weight: 2})
	edges = append(edges, &Edge{start: "x", end: "z", weight: 4})
	edges = append(edges, &Edge{start: "z", end: "s", weight: 7})
	edges = append(edges, &Edge{start: "z", end: "x", weight: 6})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.initializeGraph()

	fmt.Println("--- graph ---")
	graph.printGraph()

	fmt.Println("--- Dijkstra's algorithm ---")
	dijkstra(graph, vertices["s"])

	fmt.Println("--- result ---")
	for _, v := range vertices {
		fmt.Printf("%s (%d): ", v.name, v.distance)
		printReseult(vertices["s"], v, v.name)
		fmt.Println()
	}
}
