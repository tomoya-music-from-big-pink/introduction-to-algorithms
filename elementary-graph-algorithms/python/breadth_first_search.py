from collections import deque

from graph import Color, Edge, Graph, Vertex


def breadth_first_search(graph, s):
    s.color = Color.GRAY

    queue = deque()
    queue.append(s)

    while len(queue) > 0:
        u = queue.popleft()
        adjacencies = graph.adjacency_list[u.name]
        for v in adjacencies:
            if v.color == Color.WHITE:
                print(f"{v.name}: color change to gray.")

                v.color = Color.GRAY
                v.distance = u.distance + 1
                v.predecessor = u

                queue.append(v)

        print(f"{u.name}: color change to black.")

        u.color = Color.BLACK


def print_path(s, v, v_name):
    if v == s:
        print(s.name, end=" -> " if v.name != v_name else "")
    elif v.predecessor:
        print_path(s, v.predecessor, v_name)
        print(v.name, end=" -> " if v.name != v_name else "")


if __name__ == "__main__":
    vertices = {}
    for name in "srvwxtuy":
        vertices[name] = Vertex(name)

    edges = []
    edges.append(Edge("s", "r"))
    edges.append(Edge("s", "w"))
    edges.append(Edge("r", "s"))
    edges.append(Edge("r", "v"))
    edges.append(Edge("v", "r"))
    edges.append(Edge("w", "s"))
    edges.append(Edge("w", "t"))
    edges.append(Edge("w", "x"))
    edges.append(Edge("t", "w"))
    edges.append(Edge("t", "x"))
    edges.append(Edge("t", "u"))
    edges.append(Edge("x", "w"))
    edges.append(Edge("x", "t"))
    edges.append(Edge("x", "u"))
    edges.append(Edge("x", "y"))
    edges.append(Edge("u", "t"))
    edges.append(Edge("u", "x"))
    edges.append(Edge("u", "y"))
    edges.append(Edge("y", "x"))
    edges.append(Edge("y", "u"))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- breadth first search ---")
    breadth_first_search(graph, vertices["s"])

    print("--- result ---")
    for u in vertices.values():
        print(f"{u.name} ({u.distance}) : ", end="")
        if u != vertices["s"]:
            print_path(vertices["s"], u, u.name)
        print("")
