import heapq

from graph import Edge, Graph, Vertex, initialize_single_source, relax


def dijkstra(graph, s):
    initialize_single_source(graph, s)

    vertices = list(graph.vertices.values())

    heapq.heapify(vertices)

    adjacency_list = graph.adjacency_list
    while vertices:
        u = heapq.heappop(vertices)
        adjacencies = adjacency_list[u.name]
        for v, weight in adjacencies:
            relax(u, v, weight)

            heapq.heapify(vertices)


def print_path(s, v, v_name):
    if v == s:
        print(s.name, end=" -> " if v.name != v_name else "")
    elif v.predecessor:
        print_path(s, v.predecessor, v_name)
        print(v.name, end=" -> " if v.name != v_name else "")


if __name__ == "__main__":
    vertices = {}
    for name in "styxz":
        vertices[name] = Vertex(name)

    edges = []
    edges.append(Edge("s", "t", 10))
    edges.append(Edge("s", "y", 5))
    edges.append(Edge("t", "y", 2))
    edges.append(Edge("t", "x", 1))
    edges.append(Edge("y", "t", 3))
    edges.append(Edge("y", "x", 9))
    edges.append(Edge("y", "z", 2))
    edges.append(Edge("x", "z", 4))
    edges.append(Edge("z", "s", 7))
    edges.append(Edge("z", "x", 6))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- Dikstra's algorithm ---")
    dijkstra(graph, vertices["s"])

    print("--- result ---")
    for v in vertices.values():
        print(f"{v.name} : {v.distance} ", end="(")
        print_path(vertices["s"], v, v.name)
        print(")")
