from graph import Color, Edge, Graph, Vertex

time = 0


def depth_first_search(graph):
    vertices = graph.vertices

    for u in vertices.values():
        u.color = Color.WHITE
        u.predecessor = None

    for u in vertices.values():
        if u.color == Color.WHITE:
            depth_first_search_internal(graph, u)


def depth_first_search_internal(graph, u):
    global time

    time += 1
    u.start_time = time
    u.color = Color.GRAY
    print(f"{u.name} change color to gray. time = {time}")

    adjacencies = graph.adjacency_list[u.name]
    for v in adjacencies:
        if v.color == Color.WHITE:
            v.predecessor = u
            depth_first_search_internal(graph, v)

    time += 1
    u.finish_time = time
    u.color = Color.BLACK
    print(f"{u.name} change color to black. time = {time}")


def print_path(u, u_name):
    if not u:
        return

    print_path(u.predecessor, u_name)
    print(u.name, end=" -> " if u.name != u_name else "")


if __name__ == "__main__":
    vertices = {}
    for name in "uvxywz":
        vertices[name] = Vertex(name)

    edges = []
    edges.append(Edge("u", "v"))
    edges.append(Edge("u", "x"))
    edges.append(Edge("v", "y"))
    edges.append(Edge("x", "v"))
    edges.append(Edge("y", "x"))
    edges.append(Edge("w", "y"))
    edges.append(Edge("w", "z"))
    edges.append(Edge("z", "z"))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- depth first search ---")
    depth_first_search(graph)

    print("--- result ---")
    for u in vertices.values():
        print(f"{u.name} ({u.start_time}/{u.finish_time}) : ", end="")
        print_path(u, u.name)
        print("")
