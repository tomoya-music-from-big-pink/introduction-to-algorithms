package main

import "fmt"

func bellmanFord(g *Graph, s *Vertex) bool {
	initializeSingleSource(g, s)

	vertices := g.vertices
	edges := g.edges
	for i := range len(vertices) - 1 {
		fmt.Printf("--- i = %d ---\n", i)

		for _, edge := range edges {
			start := edge.start
			end := edge.end
			weight := edge.weight

			relax(vertices[start], vertices[end], weight)
		}
	}

	hasCycle := false
	for _, edge := range edges {
		start := edge.start
		end := edge.end
		weight := edge.weight

		if vertices[end].distance > vertices[start].distance+weight {
			hasCycle = true

			break
		}
	}

	return hasCycle
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
	edges = append(edges, &Edge{start: "s", end: "t", weight: 6})
	edges = append(edges, &Edge{start: "s", end: "y", weight: 7})
	edges = append(edges, &Edge{start: "t", end: "y", weight: 8})
	edges = append(edges, &Edge{start: "t", end: "x", weight: 5})
	edges = append(edges, &Edge{start: "t", end: "z", weight: -4})
	edges = append(edges, &Edge{start: "y", end: "x", weight: -3})
	edges = append(edges, &Edge{start: "y", end: "z", weight: 9})
	edges = append(edges, &Edge{start: "x", end: "t", weight: -2})
	edges = append(edges, &Edge{start: "z", end: "s", weight: 2})
	edges = append(edges, &Edge{start: "z", end: "x", weight: 7})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.initializeGraph()

	fmt.Println("--- graph ---")
	graph.printGraph()

	fmt.Println("--- Bellman-Ford algorithm ---")
	hasCycle := bellmanFord(graph, vertices["s"])

	fmt.Println("--- result ---")
	if hasCycle {
		fmt.Println("Graph has a cycle.")
	} else {
		for _, v := range vertices {
			fmt.Printf("%s (%d): ", v.name, v.distance)
			printReseult(vertices["s"], v, v.name)
			fmt.Println()
		}
	}
}
