from graph import Edge, Graph, Vertex, initialize_single_source, relax


def bellman_ford(graph, s):
    initialize_single_source(graph, s)

    vertices = graph.vertices
    edges = graph.edges

    for i in range(len(vertices) - 1):
        print(f"--- i = {i} ---")
        for edge in edges:
            start = edge.start
            end = edge.end
            weight = edge.weight
            relax(vertices[start], vertices[end], weight)

    has_cycle = False
    for edge in edges:
        start = edge.start
        end = edge.end
        weight = edge.weight
        if vertices[end].distance > vertices[start].distance + weight:
            has_cycle = True

            break

    return has_cycle


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
    edges.append(Edge("s", "t", 6))
    edges.append(Edge("s", "y", 7))
    edges.append(Edge("t", "y", 8))
    edges.append(Edge("t", "x", 5))
    edges.append(Edge("t", "z", -4))
    edges.append(Edge("y", "x", -3))
    edges.append(Edge("y", "z", 9))
    edges.append(Edge("x", "t", -2))
    edges.append(Edge("z", "s", 2))
    edges.append(Edge("z", "x", 7))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- Bellman-Ford algorithm ---")
    has_cycle = bellman_ford(graph, vertices["s"])

    print("--- result ---")
    if has_cycle:
        print("Graph has a cycle.")
    else:
        for v in vertices.values():
            print(f"{v.name} : {v.distance} ", end="(")
            print_path(vertices["s"], v, v.name)
            print(")")
