import enum
import sys


class Color(enum.Enum):
    WHITE = "white"
    GRAY = "gray"
    BLACK = "black"


class Vertex:
    def __init__(self, name):
        self.name = name
        self.distance = 0
        self.predecessor = None
        self.start_time = self.finish_time = 0
        self.color = Color.WHITE

    def __lt__(self, other):
        return self.distance < other.distance


class Edge:
    def __init__(self, start, end, weight=1):
        self.start = start
        self.end = end
        self.weight = weight

    def __lt__(self, other):
        return self.weight < other.weight


class Graph:
    def __init__(self, vertices, edges):
        self.vertices = vertices
        self.edges = edges
        self.adjacency_list = {}
        self.__initialize_graph()

    def print_graph(self):
        for k in self.vertices.keys():
            print(k, end=" : ")
            adjacencies = self.adjacency_list[k]
            for i, (adjacency, weight) in enumerate(adjacencies):
                print(
                    f"{adjacency.name}({weight})",
                    end=" -> " if i < len(adjacencies) - 1 else "",
                )

            print("")

    def __initialize_graph(self):
        for k, v in self.vertices.items():
            self.adjacency_list[k] = []

        for edge in self.edges:
            self.adjacency_list[edge.start].append(
                (self.vertices[edge.end], edge.weight)
            )


class UnionFind:
    def __init__(self, names):
        self.names = names
        self.parent = {name: name for name in names}
        self.rank = {name: 0 for name in names}

    def union(self, x, y):
        root_of_x = self.find_set(x)
        root_of_y = self.find_set(y)

        if root_of_x != root_of_y:
            if self.rank[root_of_x] < self.rank[root_of_y]:
                self.parent[root_of_y] = root_of_x
            elif self.rank[root_of_y] < self.rank[root_of_x]:
                self.parent[root_of_x] = root_of_y
            else:
                self.parent[root_of_y] = root_of_x
                self.rank[root_of_x] += 1

    def find_set(self, name):
        if name != self.parent[name]:
            self.parent[name] = self.find_set(self.parent[name])

        return self.parent[name]

    def print_set(self):
        groups = {}
        for x in self.names:
            root = self.find_set(x)
            groups.setdefault(root, []).append(x)
        for root, members in groups.items():
            print(f"root={root}: {members}")


def initialize_single_source(graph, s):
    vertices = graph.vertices
    for v in vertices.values():
        v.distance = sys.maxsize
        v.predecessor = None

    s.distance = 0


def relax(u, v, weight):
    if v.distance > u.distance + weight:
        print(
            f"relax: u = {u.name}, v = {v.name} / before = {v.distance}, after = {u.distance + weight}"
        )

        v.distance = u.distance + weight
        v.predecessor = u
