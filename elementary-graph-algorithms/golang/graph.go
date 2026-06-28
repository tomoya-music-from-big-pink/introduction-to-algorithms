package main

import "fmt"

type Vertex struct {
	name                                  string
	predecessor                           *Vertex
	distance, start_time, end_time, color int
}

type Edge struct {
	start, end string
}

type Graph struct {
	vertices      map[string]*Vertex
	edges         []*Edge
	adjacencyList map[string][]*Vertex
}

func (graph *Graph) initializeGraph() {
	graph.adjacencyList = make(map[string][]*Vertex)
	for _, v := range graph.vertices {
		graph.adjacencyList[v.name] = make([]*Vertex, 0)
	}
	for _, edge := range graph.edges {
		graph.adjacencyList[edge.start] = append(graph.adjacencyList[edge.start], graph.vertices[edge.end])
	}
}

func (graph *Graph) printGraph() {
	for _, v := range graph.vertices {
		fmt.Print(v.name, " : ")

		adjacencies := graph.adjacencyList[v.name]
		for i, adjacency := range adjacencies {
			fmt.Printf("%v", adjacency.name)
			if i < len(adjacencies)-1 {
				fmt.Print(" -> ")
			}
		}

		fmt.Print("\n")
	}
}
