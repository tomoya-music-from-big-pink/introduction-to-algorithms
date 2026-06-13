class Element:

    def __init__(self):
        self.keys = []
        self.n = 0
        self.is_leaf = True
        self.children = []


class Tree:

    def __init__(self, min_degree=2):
        self.min_degree = min_degree
        self.root = Element()

    def insert(self, key):
        r = self.root
        if r.n == 2 * self.min_degree - 1:
            s = Element()
            s.is_leaf = False
            self.root = s
            s.children.append(r)
            self.__split_child(s, 0)

            self.__insert_key_non_full(s, key)
        else:
            self.__insert_key_non_full(r, key)

    def __insert_key_non_full(self, x, key):
        if x.is_leaf:
            i = x.n - 1
            while i >= 0 and x.keys[i] > key:
                i -= 1

            i += 1

            x.keys.insert(i, key)
            x.n = len(x.keys)
        else:
            i = x.n - 1
            while i >= 0 and x.keys[i] > key:
                i -= 1

            i += 1

            if x.children[i].n == 2 * self.min_degree - 1:
                self.__split_child(x, i)

                if x.keys[i] < key:
                    i += 1

            self.__insert_key_non_full(x.children[i], key)

    def __split_child(self, x, i):
        y = x.children[i]
        z = Element()

        key_to_lift = y.keys[self.min_degree - 1]

        z.keys = y.keys[self.min_degree:]
        z.n = len(z.keys)
        y.keys = y.keys[:self.min_degree - 1]
        y.n = len(y.keys)

        z.is_leaf = y.is_leaf
        if not y.is_leaf:
            z.children = y.children[self.min_degree:]
            y.children = y.children[:self.min_degree]

        x.keys.insert(i, key_to_lift)
        x.n = len(x.keys)
        x.children.insert(i + 1, z)

    def print_tree(self):

        def print_tree_internal(r, space_size=0):
            if not r.keys:
                return

            print("  " * space_size, r.keys)
            if r.children:
                for i in range(len(r.children)):
                    print_tree_internal(r.children[i], space_size=space_size+1)

        print_tree_internal(self.root)


if __name__ == '__main__':
    t = Tree()

    while True:
        print("1:insert 2:remove 3:print > ", end="")
        if (op := int(input())) == 1:
            print("input key > ", end="")
            key = int(input())

            t.insert(key)
        elif op == 2:
            pass
        elif op == 3:
            t.print_tree()
        else:
            print("invalid operation")
