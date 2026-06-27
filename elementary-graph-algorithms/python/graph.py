import enum


class Color(enum.Enum):
    WHITE = "white"
    GRAY = "gray"
    BLACK = "black"


class Vertex:
    def __init__(self, name):
        self.name = name
        self.distance = 0
        self.predecessor = None
        self.color = Color.WHITE


class Edge:
    def __init__(self, start, end, weight=1):
        self.start = start
        self.end = end
        self.weight = weight


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
            for i, adjacency in enumerate(adjacencies):
                print(adjacency.name, end=" -> " if i < len(adjacencies) - 1 else "")

            print("")

    def __initialize_graph(self):
        for k, v in self.vertices.items():
            self.adjacency_list[k] = []

        for edge in self.edges:
            self.adjacency_list[edge.start].append(self.vertices[edge.end])
