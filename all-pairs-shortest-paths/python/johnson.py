from graph import Edge
from graph import Graph
from graph import Vertex
from graph import initialize_single_source
from graph import relax

import copy
import heapq
import sys


def bellman_ford(graph):
    vertices = graph.vertices
    s = vertices["0"]

    initialize_single_source(graph, s)

    edges = graph.edges
    for i in range(len(vertices) - 1):
        for edge in edges:
            u = vertices[edge.start]
            v = vertices[edge.end]
            relax(u, v, edge.weight)

    has_cycle = False
    for edge in edges:
        u = vertices[edge.start]
        v = vertices[edge.end]
        if v.distance > u.distance + edge.weight:
            has_cycle = True

            break

    return has_cycle


def dijkstra(graph, s):
    initialize_single_source(graph, s)

    vertices = list(graph.vertices.values())
    heapq.heapify(vertices)

    while len(vertices):
        u = heapq.heappop(vertices)
        adjacencies = graph.adjacency_list[u.name]
        for v, weight in adjacencies:
            relax(u, v, weight)

            heapq.heapify(vertices)

def print_path(s, v, v_name):
    if v == s:
        print(s.name, end=" -> " if v.name != v_name else "")
    elif v.predecessor:
        print_path(s, v.predecessor, v_name)
        print(v.name, end=" -> " if v.name != v_name else "")

def johnson(graph):
    new_vertices = copy.deepcopy(graph.vertices)
    new_vertices["0"] = Vertex("0")

    edges = graph.edges
    new_edges = copy.deepcopy(edges)
    for vertices in graph.vertices.values():
        edges.append(Edge("0", vertices.name, 0))
        new_edges.append(Edge("0", vertices.name, 0))

    extended_graph = Graph(new_vertices, new_edges)

    print("--- extended graph ---")
    extended_graph.print_graph()

    print("--- Bellman-Ford algorithm ---")
    has_cycle = bellman_ford(extended_graph)
    if has_cycle:
        print("Extended graph has a cycle")

        return

    print("--- calculate h ---")
    h = {}
    for vertex in new_vertices.values():
        h[vertex.name] = vertex.distance

    for k, v in h.items():
        print(f"vertex:{k}, value = {v}")

    edges_dict = {}
    for edge in graph.edges:
        edges_dict[f"{edge.start}-{edge.end}"] = edge
    new_edges_dict = {}
    for new_edge in new_edges:
        new_edges_dict[f"{new_edge.start}-{new_edge.end}"] = new_edge

    print("--- re-calculate weight ---")
    for new_edge in new_edges:
        u = new_edge.start
        v = new_edge.end
        new_edges_dict[f"{u}-{v}"].weight = edges_dict[f"{u}-{v}"].weight + h[u] - h[v]

    for edge in edges:
        print(f"start = {edge.start}, end = {edge.end}, weight = {edge.weight}")

    D = [[-1 for _ in range(len(graph.vertices))] for _ in range(len(graph.vertices))]
    print("--- Dijstra's algorithm ---")
    for s in graph.vertices.values():
        print(f"start = {s.name}")
        dijkstra(graph, s)

        for v in graph.vertices.values():
            d = v.distance + h[s.name] - h[v.name]
            D[int(s.name) - 1][int(v.name) - 1] = d
            print(f"{v.name} : {d} / {v.distance} ", end="(")
            print_path(s, v, v.name)
            print(")")

    return D


def print_matrix(A):
    if not A:
        return

    for row in A:
        for i in range(len(row)):
            print(f"{row[i]:3d}" if row[i] < sys.maxsize else "INF", end=", ")
        print("")


if __name__ == '__main__':
    vertices = {}
    for name in "12345":
        vertices[name] = Vertex(name)

    edges = []
    edges.append(Edge("1", "2", 3))
    edges.append(Edge("1", "3", 8))
    edges.append(Edge("1", "5", -4))
    edges.append(Edge("2", "4", 1))
    edges.append(Edge("2", "5", 7))
    edges.append(Edge("3", "2", 4))
    edges.append(Edge("4", "1", 2))
    edges.append(Edge("4", "3", -5))
    edges.append(Edge("5", "4", 6))

    graph = Graph(vertices, edges)

    print("--- graph ---")
    graph.print_graph()

    print("--- Johnson's algorithm ---")
    D = johnson(graph)

    print("--- result ---")
    print_matrix(D)
