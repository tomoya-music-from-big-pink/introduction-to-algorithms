from ast import arguments
import enum

from collections import defaultdict
from collections import deque
import sys


class Color(enum.Enum):

    BLACK = 'black'
    GRAY = 'gray'
    WHITE = 'white'


class Vertex:

    def __init__(self, name, color=Color.WHITE):
        self.name = name
        self.color = color
        self.predecessor = None


class Edge:

    def __init__(self, start, end, capacity):
        self.start = start
        self.end = end
        self.flow = 0
        self.capacity = capacity

    @property
    def residual_capacity(self):
        return self.capacity - self.flow

    def __repr__(self):
        return f"{self.start} -> {self.end} ({self.residual_capacity})"

class Graph:

    def __init__(self, vertices, edges):
        self.vertices = vertices
        self.edges = edges
        self.edges_dict = {}
        self.adjacency_list = defaultdict(list)
        self.residual_network = defaultdict(list)

        self.__initialize_graph()

    def print_graph(self):
        self.__print_graph_internal()

    def print_final_result(self):
        self.__print_graph_internal(finished=True)

    def __print_graph_internal(self, finished=False):
        if not finished:
            print("source: sink (flow / capacity / residual-capacity)")
        else:
            print("source: sink (flow)")

        total = 0
        for u in self.vertices.values():
            print(u.name, ': ', end='')
            adjacencies = self.adjacency_list[u.name]
            for i, (v, edge) in enumerate(adjacencies):
                print(f"{v.name} ({edge.flow}", end='')
                if not finished:
                    print(f"/ {edge.capacity} / {edge.residual_capacity}", end='')
                print(")", end=" -> " if i < len(adjacencies) - 1 else '')

                if finished and v.name == "t":
                    total += edge.flow

            print()

        print(f"total flow = {total}")

    def construct_residual_network(self):
        self.residual_network.clear()

        for edge in self.edges:
            start = edge.start
            end = edge.end
            flow = edge.flow
            residual_capacity = edge.residual_capacity

            if flow > 0:
                self.residual_network[end].append((self.vertices[start], flow))
            if residual_capacity > 0:
                self.residual_network[start].append((self.vertices[end], residual_capacity))

    def breadth_first_search(self, s):
        for u in self.vertices.values():
            u.color = Color.WHITE
            u.predecessor = None

        s.color = Color.GRAY

        q = deque([s])

        while len(q) > 0:
            u = deque.popleft(q)

            adjacencies = self.residual_network[u.name]
            for adj, _ in adjacencies:

                v = self.vertices[adj.name]
                if v.color == Color.WHITE:
                    v.color = Color.GRAY
                    v.predecessor = u
                    q.append(v)

            u.color = Color.BLACK

    def find_aumenting_path(self):
        exists_aumenting_path = False
        new_flow = sys.maxsize

        v = self.vertices['t']
        aumenting_paths = []
        while v.predecessor:
            predecessor = v.predecessor
            edge = self.edges_dict[f"{predecessor.name}-{v.name}"]
            new_flow = min(edge.residual_capacity, new_flow)
            aumenting_paths.append(edge)
            v = predecessor

            if v.name == 's':
                exists_aumenting_path = True

        return exists_aumenting_path, new_flow, aumenting_paths[::-1]

    def print_residual_network(self):
        for u in self.vertices.values():
            print(u.name, ':', end='')
            adjacencies = self.residual_network[u.name]
            for i, (v, weight) in enumerate(adjacencies):
                print(f"{v.name} ({weight})", end=' -> ' if i < len(adjacencies) - 1 else '')
            print()

    def __initialize_graph(self):
        for edge in self.edges:
            start = edge.start
            end = edge.end
            self.edges_dict[f"{start}-{end}"] = edge
            v = self.vertices[end]
            self.adjacency_list[start].append((v, edge))
