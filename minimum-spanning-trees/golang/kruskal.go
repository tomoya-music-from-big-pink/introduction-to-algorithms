package main

import (
	"fmt"
	"sort"
	"strings"
)

func krsukal(graph *Graph) []*Edge {
	edges := graph.edges

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	var names []string
	for name, _ := range graph.vertices {
		names = append(names, name)
	}

	unionFind := &UnionFind{}
	unionFind.initializeUnionSet(strings.Join(names, ""))

	var results []*Edge
	for _, edge := range edges {
		start := edge.start
		end := edge.end
		if unionFind.findSet(start) != unionFind.findSet(end) {
			fmt.Printf("append edge: start = %s, end = %s, weight = %d\n", start, end, edge.weight)
			results = append(results, edge)

			unionFind.union(start, end)
		}
	}

	return results
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

	fmt.Println("--- Kruskal's algorithm ---")
	results := krsukal(graph)

	fmt.Println("--- result ---")
	total := 0
	for _, edge := range results {
		fmt.Printf("start = %s, end = %s, weight = %d\n", edge.start, edge.end, edge.weight)

		total += edge.weight
	}
	fmt.Printf("total = %d\n", total)
}
