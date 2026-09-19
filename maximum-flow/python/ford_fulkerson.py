from graph import Edge
from graph import Graph
from graph import Vertex


def ford_fulkerson(graph):
    vertices = graph.vertices

    i = 1
    while True:
        graph.construct_residual_network()

        graph.breadth_first_search(vertices['s'])

        exists_aumenting_path, new_flow, aumenting_path= graph.find_aumenting_path()
        if not exists_aumenting_path:
            break

        print(f"--- i = {i} ---")

        print("--- augmenting path ---")
        print(aumenting_path)

        for edge in aumenting_path:
            edge.flow += new_flow

        print("--- graph after updated flow ---")
        graph.print_graph()

        if i == 4:
            break

        i += 1


if __name__ == '__main__':
    vertices = {}
    for name in ['s', 'v1', 'v2', 'v3', 'v4', 't']:
        vertices[name] = Vertex(name)

    edges = []
    edges.append(Edge('s', 'v1', 16))
    edges.append(Edge('s', 'v2', 13))
    edges.append(Edge('v1', 'v3', 12))
    edges.append(Edge('v2', 'v1', 4))
    edges.append(Edge('v2', 'v4', 14))
    edges.append(Edge('v3', 'v2', 9))
    edges.append(Edge('v3', 't', 20))
    edges.append(Edge('v4', 'v3', 7))
    edges.append(Edge('v4', 't', 4))

    graph = Graph(vertices, edges)

    print("--- Graph ---")
    graph.print_graph()

    print("--- Ford-Fulkerson method ---")
    ford_fulkerson(graph)

    print("--- final result ---")
    graph.print_final_result()
