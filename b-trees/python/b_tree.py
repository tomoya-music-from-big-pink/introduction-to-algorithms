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

    def contains(self, key):

        def contains_internal(x, key):
            i = 0
            while i < x.n and x.keys[i] < key:
                i += 1

            if i < x.n and x.keys[i] == key:
                return True
            else:
                if x.is_leaf:
                    return False
                else:
                    return contains_internal(x.children[i], key)

        return contains_internal(self.root, key)

    def remove(self, key):
        x = self.contains(key)
        if not x:
            print(f"{key} is not found in the tree.")

        self.__remove_internal(self.root, key)

    def __remove_internal(self, x, key):
        i = 0
        while i < x.n and x.keys[i] < key:
            i += 1

        if i < x.n and x.keys[i] == key:
            if x.is_leaf:
                x.keys.pop(i)
                x.n -= 1
            else:
                if x.children[i].n >= self.min_degree:
                    y = x.children[i]
                    key_to_lift = y.keys[-1]

                    y.keys = y.keys[:-1]
                    y.keys.append(x.keys[i])
                    x.keys[i] = key_to_lift

                    self.__remove_internal(x.children[i], key)
                elif x.children[i + 1].n >= self.min_degree:
                    y = x.children[i + 1]
                    key_to_lift = y.keys[0]

                    y.keys = y.keys[1:]
                    y.keys.insert(0, x.keys[i])
                    x.keys[i] = key_to_lift

                    self.__remove_internal(x.children[i + 1], key)
                else:
                    y = x.children[i]
                    z = x.children[i + 1]

                    y.keys.append(x.keys[i])
                    y.keys.extend(z.keys)
                    y.n = len(y.keys)

                    x.keys.pop(i)
                    x.children.pop(i)
                    x.n -= 1
                    if x.n == 0:
                        self.root = y

                    self.__remove_internal(y, key)
        else:
            if x.children[i].n == self.min_degree - 1:
                if i > 0 and x.children[i - 1].n >= self.min_degree:
                    y = x.children[i]
                    z = x.children[i - 1]
                    key_to_lift = z.keys[-1]
                    key_to_drop = x.keys[i - 1]

                    z.keys = z.keys[:-1]
                    z.n -= 1
                    y.keys.insert(0, key_to_drop)
                    y.n += 1
                    x.keys[i - 1] = key_to_lift

                    if not y.is_leaf:
                        children_to_move = z.children[-1]
                        y.children.insert(0, children_to_move)
                        z.children = z.children[:-1]
                elif i < x.n and x.children[i + 1].n >= self.min_degree:
                    y = x.children[i]
                    z = x.children[i + 1]
                    key_to_lift = z.keys[0]
                    key_to_drop = x.keys[i]

                    z.keys = z.keys[1:]
                    z.n -= 1
                    y.keys.append(key_to_drop)
                    y.n += 1
                    x.keys[i] = key_to_lift

                    if not y.is_leaf:
                        children_to_move = z.children[0]
                        y.children.append(children_to_move)
                        z.children = z.children[1:]

                self.__remove_internal(x.children[i], key)
            else:
                self.__remove_internal(x.children[i], key)

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

        z.keys = y.keys[self.min_degree :]
        z.n = len(z.keys)
        y.keys = y.keys[: self.min_degree - 1]
        y.n = len(y.keys)

        z.is_leaf = y.is_leaf
        if not y.is_leaf:
            z.children = y.children[self.min_degree :]
            y.children = y.children[: self.min_degree]

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
                    print_tree_internal(r.children[i], space_size=space_size + 1)

        print_tree_internal(self.root)


if __name__ == "__main__":
    t = Tree()

    t.insert(5)
    t.insert(9)
    t.insert(3)
    t.insert(7)
    # t.insert(1)
    # t.insert(2)
    # t.insert(8)
    # t.insert(6)
    # t.insert(0)
    # t.insert(4)

    while True:
        print("1:insert 2:remove 3:print > ", end="")
        if (op := int(input())) == 1:
            print("input key > ", end="")
            key = int(input())

            t.insert(key)
        elif op == 2:
            print("input key > ", end="")
            key = int(input())

            t.remove(key)
        elif op == 3:
            t.print_tree()
        else:
            print("invalid operation")
