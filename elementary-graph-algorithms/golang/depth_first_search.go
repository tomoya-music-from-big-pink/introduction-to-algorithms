package main

import "fmt"

var time = 0

func depthFirstSearch(graph *Graph) {
	vertices := graph.vertices
	for _, u := range vertices {
		u.predecessor = nil
		u.color = 0
		u.start_time = 0
		u.end_time = 0
	}

	for _, u := range vertices {
		if u.color == 0 {
			depthFirstSearchInternal(graph, u)
		}
	}
}

func depthFirstSearchInternal(graph *Graph, u *Vertex) {
	time += 1
	fmt.Printf("%s: change color to gray. time = %d\n", u.name, time)
	u.start_time = time
	u.color = 1

	for _, v := range graph.adjacencyList[u.name] {
		if v.color == 0 {
			v.predecessor = u
			depthFirstSearchInternal(graph, v)
		}
	}

	time += 1
	fmt.Printf("%s: change color to black. time = %d\n", u.name, time)
	u.end_time = time
	u.color = 2
}

func printPath(v *Vertex, name string) {
	if v == nil {
		return
	}

	printPath(v.predecessor, name)
	fmt.Printf("%s", v.name)
	if v.name != name {
		fmt.Print(" -> ")
	}
}

func main() {
	vertices := make(map[string]*Vertex, 6)
	for _, name := range "uvxywz" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0, 8)
	edges = append(edges, &Edge{start: "u", end: "v"})
	edges = append(edges, &Edge{start: "u", end: "x"})
	edges = append(edges, &Edge{start: "v", end: "y"})
	edges = append(edges, &Edge{start: "x", end: "v"})
	edges = append(edges, &Edge{start: "y", end: "x"})
	edges = append(edges, &Edge{start: "w", end: "y"})
	edges = append(edges, &Edge{start: "w", end: "z"})
	edges = append(edges, &Edge{start: "z", end: "z"})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.initializeGraph()

	fmt.Println("--- graph ---")
	graph.printGraph()

	fmt.Println("--- depth first search ---")
	depthFirstSearch(graph)

	fmt.Println("--- result ---")
	for _, v := range vertices {
		fmt.Printf("%s (%d/%d): ", v.name, v.start_time, v.end_time)
		printPath(v, v.name)
		fmt.Print("\n")
	}
}
