package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
)

func johnson(graph *Graph) [][]int {
	vertices := graph.vertices
	edges := graph.edges

	extendedVertices := make(map[string]*Vertex, len(vertices)+1)
	extendedVertices["0"] = &Vertex{name: string("0")}
	for k, v := range vertices {
		extendedVertices[k] = v
	}

	extendedEdges := make([]*Edge, 0)
	for k, _ := range vertices {
		extendedEdges = append(extendedEdges, &Edge{"0", k, 0})
	}
	for _, edge := range edges {
		extendedEdges = append(extendedEdges, edge)
	}

	extended_graph := &Graph{vertices: extendedVertices, edges: extendedEdges}
	extended_graph.initializeGraph()

	fmt.Println("--- extended graph ---")
	extended_graph.printGraph()

	fmt.Println("--- Bellman-Ford algorithm ---")
	hasCycle := bellmanFord(extended_graph, extendedVertices["0"])
	if hasCycle {
		fmt.Println("Graph has a cycle.")

		return nil
	}

	fmt.Println("--- calculate h ---")
	h := make(map[string]int)
	for k, v := range extendedVertices {
		h[k] = v.distance
		fmt.Printf("%s: %d\n", k, h[k])
	}

	fmt.Println("--- re-calculate weight ---")
	weightDict := make(map[string]int)
	for k, _ := range vertices {
		weightDict[fmt.Sprint("0-%s", k)] = 0
	}
	for _, edge := range edges {
		weightDict[fmt.Sprintf("%s-%s", edge.start, edge.end)] = edge.weight

	}
	extendedWeightDict := make(map[string]int)
	for _, extendedEdge := range extendedEdges {
		targetWeight := weightDict[fmt.Sprintf("%s-%s", extendedEdge.start, extendedEdge.end)]
		extendedWeightDict[fmt.Sprintf("%s-%s", extendedEdge.start, extendedEdge.end)] = targetWeight + h[extendedEdge.start] - h[extendedEdge.end]
	}
	for k, v := range extendedWeightDict {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("--- Dikstra's algorithm ---")
	D := make([][]int, len(vertices))
	for i := range len(D) {
		D[i] = make([]int, len(vertices))
	}
	for _, s := range vertices {
		fmt.Printf("start = %s\n", s.name)
		dijkstra(graph, s)

		for _, v := range vertices {
			d := v.distance + h[s.name] - h[v.name]
			i, _ := strconv.Atoi(s.name)
			j, _ := strconv.Atoi(v.name)
			D[i-1][j-1] = d
			fmt.Printf("%s (%d/%d): ", v.name, d, v.distance)
			printReseult(s, v, v.name)
			fmt.Println()
		}
	}

	return D
}

func bellmanFord(graph *Graph, s *Vertex) bool {
	initializeSingleSource(graph, s)

	vertices := graph.vertices
	edges := graph.edges

	for range len(vertices) - 1 {
		for _, edge := range edges {
			start := edge.start
			end := edge.end
			weight := edge.weight
			relax(vertices[start], vertices[end], weight)
		}
	}

	hasCycle := false
	for _, edge := range edges {
		u := vertices[edge.start]
		v := vertices[edge.end]
		weight := edge.weight
		if v.distance > u.distance+weight {
			hasCycle = true

			break
		}
	}

	return hasCycle
}

func dijkstra(graph *Graph, s *Vertex) {
	vertexHeap := make(vertexHeap, 0)
	vertices := graph.vertices

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

	adjacencyList := graph.adjacencyList
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

func printMatrix(W [][]int) {
	for i := range W {
		for j := range W[i] {
			if W[i][j] != math.MaxInt8 {
				fmt.Printf("%3d", W[i][j])
			} else {
				fmt.Print("  -")
			}
		}

		fmt.Println("")
	}
}

func main() {
	vertices := make(map[string]*Vertex)
	for _, name := range "12345" {
		vertices[string(name)] = &Vertex{name: string(name)}
	}

	edges := make([]*Edge, 0)
	edges = append(edges, &Edge{"1", "2", 3})
	edges = append(edges, &Edge{"1", "3", 8})
	edges = append(edges, &Edge{"1", "5", -4})
	edges = append(edges, &Edge{"2", "4", 1})
	edges = append(edges, &Edge{"2", "5", 7})
	edges = append(edges, &Edge{"3", "2", 4})
	edges = append(edges, &Edge{"4", "1", 2})
	edges = append(edges, &Edge{"4", "3", -5})
	edges = append(edges, &Edge{"5", "4", 6})

	graph := &Graph{vertices: vertices, edges: edges}
	graph.initializeGraph()

	fmt.Println("--- graph ---")
	graph.printGraph()

	fmt.Println("--- Johnson's algorithm ---")
	D := johnson(graph)

	fmt.Println("--- result ---")
	printMatrix(D)
}
