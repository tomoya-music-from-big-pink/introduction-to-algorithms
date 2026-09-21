package main

import (
	"fmt"
	"math"
	"slices"
)

type Color int

const (
	WHITE Color = iota
	GRAY
	BLACK
)

type Vertex struct {
	color       Color
	name        string
	predecessor *Vertex
}

type Edge struct {
	start, end                       string
	flow, capacity, residualCapacity int
}

type Graph struct {
	vertices        map[string]*Vertex
	edges           []*Edge
	edgesMap        map[string]*Edge
	adjacencyList   map[*Vertex][]*Edge
	residualNetwork map[*Vertex][]*Edge
}

func (graph *Graph) InitializeGraph() {
	vertices := graph.vertices
	edges := graph.edges

	graph.edgesMap = make(map[string]*Edge)
	graph.adjacencyList = make(map[*Vertex][]*Edge)
	graph.residualNetwork = make(map[*Vertex][]*Edge)

	for _, edge := range edges {
		start := edge.start
		end := edge.end
		graph.edgesMap[fmt.Sprintf("%s-%s", start, end)] = edge

		v := vertices[edge.start]
		graph.adjacencyList[v] = append(graph.adjacencyList[v], edge)
	}
}

func (graph *Graph) PrintGraph() {
	graph.printGraphInternal(true)
}

func (graph *Graph) PrintResidualNetwork() {
	graph.printGraphInternal(false)
}

func (graph *Graph) printGraphInternal(isAdjacency bool) {
	vertices := graph.vertices

	fmt.Println("--- format ---")
	if isAdjacency {
		fmt.Println("u : v (flow / capacity / residual-capacity) -> ...")
	} else {
		fmt.Println("u : v (flow / residual-capacity) -> ...")
	}
	fmt.Println("---")

	for _, vertex := range vertices {
		fmt.Printf("%s : ", vertex.name)

		var adjacencies []*Edge
		if isAdjacency {
			adjacencies = graph.adjacencyList[vertex]
		} else {
			adjacencies = graph.residualNetwork[vertex]
		}
		for i, adjacency := range adjacencies {
			if isAdjacency {
				fmt.Printf("%s (%d / %d / %d)", adjacency.end, adjacency.flow, adjacency.capacity, adjacency.residualCapacity)
			} else {
				fmt.Printf("%s (%d / %d)", adjacency.end, adjacency.flow, adjacency.residualCapacity)
			}
			if i < len(adjacencies)-1 {
				fmt.Print(" -> ")
			}
		}

		fmt.Println()
	}
}

func (graph *Graph) ConstructResidualNetwork() {
	clear(graph.residualNetwork)

	vertices := graph.vertices

	for _, edge := range graph.edges {
		start := edge.start
		end := edge.end

		if edge.flow > 0 {
			edge.residualCapacity = edge.capacity - edge.flow
			graph.residualNetwork[vertices[end]] = append(graph.residualNetwork[vertices[end]], edge)
		}
		if edge.residualCapacity > 0 {
			graph.residualNetwork[vertices[start]] = append(graph.residualNetwork[vertices[start]], edge)
		}
	}
}

func (graph *Graph) BreadthFirstSearch() {
	vertices := graph.vertices

	for _, vertex := range vertices {
		vertex.color = WHITE
		vertex.predecessor = nil
	}

	s := vertices["s"]
	s.color = GRAY
	fmt.Printf("%s color changes to gray.\n", s.name)

	q := make([]*Vertex, 0)
	q = append(q, s)

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		fmt.Printf("target vertex: %s\n", u.name)

		adjacencies := graph.residualNetwork[u]
		for _, adjacency := range adjacencies {
			v := vertices[adjacency.end]
			if v.color == WHITE {
				fmt.Printf("%s color changes to gray.\n", v.name)

				v.color = GRAY
				v.predecessor = u

				q = append(q, v)
			}
		}

		u.color = BLACK
		fmt.Printf("%s color changes to black.\n", u.name)
	}
}

func (graph *Graph) hasAugmentingPath() (bool, []*Edge, int) {
	vertices := graph.vertices
	edges := graph.edgesMap

	p := vertices["t"]
	hasPath := false
	augmentingPathEdge := make([]*Edge, 0)
	newFlow := math.MaxFloat64

	for p.predecessor != nil {
		q := p.predecessor
		augmentingPathEdge = append(augmentingPathEdge, edges[fmt.Sprintf("%s-%s", q.name, p.name)])
		edge := edges[fmt.Sprintf("%s-%s", q.name, p.name)]
		newFlow = math.Min(float64(newFlow), float64(edge.residualCapacity))
		if q.name == "s" {
			hasPath = true
			break
		}

		p = q
	}

	slices.Reverse(augmentingPathEdge)

	return hasPath, augmentingPathEdge, int(newFlow)
}
