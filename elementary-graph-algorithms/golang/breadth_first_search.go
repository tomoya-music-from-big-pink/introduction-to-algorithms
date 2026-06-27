package main

import "fmt"

func breadthFirstSearch(graph *Graph, s *Vertex) {
	queue := make([]*Vertex, 0)
	queue = append(queue, s)

	for len(queue) > 0 {
		u := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		adjacencies := graph.adjacencyList[u.name]
		for _, v := range adjacencies {
			if v.color == 0 {
				fmt.Printf("%s: color changes to gray.\n", v.name)

				v.color = 1
				v.distance = u.distance + 1
				v.predecessor = u

				queue = append(queue, v)
			}
		}

		u.color = 2

		fmt.Printf("%s: color changes to black.\n", u.name)
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
	for _, name := range "srvwxtuy" {
		vertices[string(name)] = &Vertex{name: string(name), predecessor: nil, distance: 0}
	}

	edges := make([]*Edge, 0, 8)
	edges = append(edges, &Edge{start: "s", end: "r"})
	edges = append(edges, &Edge{start: "s", end: "w"})
	edges = append(edges, &Edge{start: "r", end: "s"})
	edges = append(edges, &Edge{start: "r", end: "v"})
	edges = append(edges, &Edge{start: "v", end: "r"})
	edges = append(edges, &Edge{start: "w", end: "s"})
	edges = append(edges, &Edge{start: "w", end: "t"})
	edges = append(edges, &Edge{start: "w", end: "x"})
	edges = append(edges, &Edge{start: "t", end: "w"})
	edges = append(edges, &Edge{start: "t", end: "x"})
	edges = append(edges, &Edge{start: "t", end: "u"})
	edges = append(edges, &Edge{start: "x", end: "w"})
	edges = append(edges, &Edge{start: "x", end: "t"})
	edges = append(edges, &Edge{start: "x", end: "u"})
	edges = append(edges, &Edge{start: "x", end: "y"})
	edges = append(edges, &Edge{start: "u", end: "t"})
	edges = append(edges, &Edge{start: "u", end: "x"})
	edges = append(edges, &Edge{start: "u", end: "y"})
	edges = append(edges, &Edge{start: "y", end: "x"})
	edges = append(edges, &Edge{start: "y", end: "u"})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.initializeGraph()

	fmt.Println("--- graph ---")
	graph.printGraph()

	fmt.Println("--- breadth first search ---")
	breadthFirstSearch(graph, vertices["s"])

	for _, v := range vertices {
		fmt.Printf("%s (%d) : ", v.name, v.distance)
		printReseult(vertices["s"], v, v.name)
		fmt.Print("\n")
	}
}
