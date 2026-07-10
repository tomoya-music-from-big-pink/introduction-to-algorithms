from graph import Edge, Graph, UnionFind, Vertex


def kruskal(graph):
    vertices = graph.vertices
    edges = graph.edges
    union_find = UnionFind("".join(vertices.keys()))

    edges.sort()

    results = []
    for edge in edges:
        start = edge.start
        end = edge.end
        weight = edge.weight
        if union_find.find_set(start) != union_find.find_set(end):
            print(f"append edge: start = {start}, end = {end}, weight = {weight}")

            results.append(edge)

            union_find.union(start, end)

    return results


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

    print("--- Kruskal's algorithm ---")
    results = kruskal(graph)

    print("--- result ---")
    total = 0
    for edge in results:
        print(f"start = {edge.start}, end = {edge.end}, weight = {edge.weight}")

        total += edge.weight

    print(f"total = {total}")
