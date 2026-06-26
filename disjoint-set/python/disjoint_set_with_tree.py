class UnionFind:
    def __init__(self, size):
        self.parent = [i for i in range(size)]
        self.rank = [1] * size

    def union(self, x, y):
        root_of_x = self.find_set(x)
        root_of_y = self.find_set(y)

        if root_of_x != root_of_y:
            if self.rank[root_of_x] > self.rank[root_of_y]:
                self.parent[root_of_y] = root_of_x
            elif self.rank[root_of_y] > self.rank[root_of_x]:
                self.parent[root_of_x] = root_of_y
            else:
                self.parent[root_of_y] = root_of_x
                self.rank[root_of_x] += 1

    def find_set(self, x):
        if x != self.parent[x]:
            self.parent[x] = self.find_set(self.parent[x])

        return self.parent[x]

    def print_set(self):
        groups = {}
        for i in range(len(self.parent)):
            root = self.find_set(i)
            groups.setdefault(root, []).append(i)
        for root, members in groups.items():
            print(f"root={root}: {members}")


if __name__ == "__main__":
    disjoint_set = UnionFind(10)

    print("--- after initialize ---")
    disjoint_set.print_set()

    print("union 1 and 3")
    disjoint_set.union(1, 3)
    disjoint_set.print_set()
    print("union 4 and 6")
    disjoint_set.union(4, 6)
    disjoint_set.print_set()
    print("union 0 and 2")
    disjoint_set.union(0, 2)
    disjoint_set.print_set()
    print("union 7 and 8")
    disjoint_set.union(7, 8)
    disjoint_set.print_set()
    print("union 0 and 1")
    disjoint_set.union(0, 1)
    disjoint_set.print_set()
    print("union 4 and 5")
    disjoint_set.union(4, 5)
    disjoint_set.print_set()
    print("union 1 and 2")
    disjoint_set.union(1, 2)
    disjoint_set.print_set()
