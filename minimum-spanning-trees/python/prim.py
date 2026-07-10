import heapq
import sys

from graph import Edge, Graph, Vertex


def prim(graph, r):
    vertices = list(graph.vertices.values())
    adjacency_list = graph.adjacency_list

    for u in vertices:
        u.distance = sys.maxsize
        u.predecessor = None

    r.distance = 0

    heapq.heapify(vertices)

    while len(vertices):
        u = heapq.heappop(vertices)
        adjacencies = adjacency_list[u.name]
        for v, weight in adjacencies:
            if v in vertices and weight < v.distance:
                print(
                    f"u = {u.name}, v = {v.name} / before = {v.distance}, after = {weight}"
                )

                v.distance = weight
                v.predecessor = u

                heapq.heapify(vertices)


if __name__ == "__main__":
    vertices = {}
    for name in "abcdefghi":
        vertices[name] = Vertex(name)

    edges = []
    edges.append(Edge("a", "b", 4))
    edges.append(Edge("a", "h", 8))
    edges.append(Edge("b", "a", 4))
    edges.append(Edge("b", "c", 8))
    edges.append(Edge("b", "h", 11))
    edges.append(Edge("c", "b", 8))
    edges.append(Edge("c", "d", 7))
    edges.append(Edge("c", "f", 4))
    edges.append(Edge("c", "i", 2))
    edges.append(Edge("d", "c", 7))
    edges.append(Edge("d", "e", 9))
    edges.append(Edge("d", "f", 14))
    edges.append(Edge("e", "d", 9))
    edges.append(Edge("e", "f", 10))
    edges.append(Edge("f", "c", 4))
    edges.append(Edge("f", "d", 14))
    edges.append(Edge("f", "e", 10))
    edges.append(Edge("f", "g", 2))
    edges.append(Edge("g", "f", 2))
    edges.append(Edge("g", "h", 1))
    edges.append(Edge("g", "i", 6))
    edges.append(Edge("h", "a", 8))
    edges.append(Edge("h", "b", 11))
    edges.append(Edge("h", "g", 1))
    edges.append(Edge("h", "i", 7))
    edges.append(Edge("i", "c", 2))
    edges.append(Edge("i", "g", 6))
    edges.append(Edge("i", "h", 7))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- Prim's algorithm ---")
    prim(graph, vertices["a"])

    print("--- result ---")
    for u in vertices.values():
        print(
            f"vertices = {u.name}, predecessor = {'None' if not u.predecessor else u.predecessor.name}, weight = {u.distance}"
        )
