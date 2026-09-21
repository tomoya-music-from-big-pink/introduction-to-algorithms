package main

import "fmt"

func printAugmentPath(edges []*Edge) {
	for i, edge := range edges {
		fmt.Printf("%s -> %s (%d)", edge.start, edge.end, edge.residualCapacity)
		if i < len(edges)-1 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

func updateFlow(augmentingPathEdge []*Edge, newFlow int) {
	for _, edge := range augmentingPathEdge {
		fmt.Printf("%s -> %s: flow : %d -> %d, residual capacity: %d -> %d\n", edge.start, edge.end, edge.flow, newFlow, edge.residualCapacity, edge.capacity-newFlow)

		edge.flow += newFlow
		edge.residualCapacity = edge.capacity - edge.flow
	}
}

func fordFulkerson(graph *Graph) {
	i := 1
	for {
		fmt.Printf("--- i = %d ---\n", i)

		fmt.Println("--- Construct residual-network ---")
		graph.ConstructResidualNetwork()
		graph.PrintResidualNetwork()

		fmt.Println("--- Bread First Search ---")
		graph.BreadthFirstSearch()

		hasPath, augmentingPathEdge, newFlow := graph.hasAugmentingPath()
		if !hasPath {
			fmt.Println("augmenting path is not found.")

			break
		}

		fmt.Println("--- Augmenting path ---")
		printAugmentPath(augmentingPathEdge)
		fmt.Println("--- new flow ---")
		fmt.Println(newFlow)

		fmt.Println("--- Update flow ---")
		updateFlow(augmentingPathEdge, newFlow)

		fmt.Println("--- graph after updating flow ---")
		graph.PrintGraph()

		i++
	}
}

func main() {
	vertices := make(map[string]*Vertex)
	vertices["s"] = &Vertex{name: "s"}
	vertices["v1"] = &Vertex{name: "v1"}
	vertices["v2"] = &Vertex{name: "v2"}
	vertices["v3"] = &Vertex{name: "v3"}
	vertices["v4"] = &Vertex{name: "v4"}
	vertices["t"] = &Vertex{name: "t"}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{start: "s", end: "v1", capacity: 16, residualCapacity: 16})
	edges = append(edges, &Edge{start: "s", end: "v2", capacity: 13, residualCapacity: 13})
	edges = append(edges, &Edge{start: "v1", end: "v3", capacity: 12, residualCapacity: 12})
	edges = append(edges, &Edge{start: "v2", end: "v1", capacity: 4, residualCapacity: 4})
	edges = append(edges, &Edge{start: "v2", end: "v4", capacity: 14, residualCapacity: 14})
	edges = append(edges, &Edge{start: "v3", end: "v2", capacity: 9, residualCapacity: 9})
	edges = append(edges, &Edge{start: "v3", end: "t", capacity: 20, residualCapacity: 20})
	edges = append(edges, &Edge{start: "v4", end: "v3", capacity: 7, residualCapacity: 7})
	edges = append(edges, &Edge{start: "v4", end: "t", capacity: 4, residualCapacity: 4})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.InitializeGraph()

	fmt.Println("--- graph ---")
	graph.PrintGraph()

	fmt.Println("--- Fork-Fulkerson method ---")
	fordFulkerson(graph)

	fmt.Println("--- final result ---")
	graph.PrintGraph()
	totalFlow := 0
	for _, edge := range graph.edges {
		if edge.end == "t" {
			totalFlow += edge.flow
		}
	}
	fmt.Printf("total flow = %d\n", totalFlow)
}
