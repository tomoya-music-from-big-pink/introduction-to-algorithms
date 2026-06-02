import enum


class Color(enum.Enum):
    RED = "red"
    BLACK = "black"


class Element:
    __slots__ = ["key", "left", "right", "parent", "color"]

    def __init__(self, key, color=Color.RED):
        self.key = key
        self.left = self.right = self.parent = None
        self.color = color

    def __str__(self):
        return f"{self.key} ({self.color.value})"


sentinel = None


def nil():
    global sentinel

    if not sentinel:
        sentinel = Element(-1, color=Color.BLACK)

    return sentinel


class Tree:
    __slots__ = ["size", "height", "root"]

    def __init__(self):
        self.root = nil()
        self.size = self.height = 0

    def put(self, key):
        self.size += 1

        z = Element(key)
        x, y = self.root, nil()
        while x != nil():
            y = x
            x = x.left if z.key < x.key else x.right

        z.parent = y
        z.left = z.right = nil()
        if y == nil():
            self.root = z
        else:
            if z.key < y.key:
                y.left = z
            else:
                y.right = z

        self.__fix_after_insertion(z)

        self.height = self.__get_height()

    def contains(self, key):
        p = self.root
        while p and p.key != key:
            p = p.left if key < p.key else p.right

        return p

    def remove(self, key):
        z = self.contains(key)
        if not z:
            print(f"{key} is not found in the tree")

            return

        self.size -= 1

        if not z.left:
            self.__transplant(z, z.right)
        elif not z.right:
            self.__transplant(z, z.left)
        else:
            y = self.__successor(z)
            if y.parent != z:
                self.__transplant(y, y.right)
                y.right = z.right
                y.right.parent = y

            self.__transplant(z, y)
            y.left = z.left
            y.left.parent = y

        self.height = self.__get_height()

    def print_tree(self):

        def print_tree_internal(r, space_size=0):
            if r.left != nil():
                print_tree_internal(r.left, space_size=space_size + 1)

            print("    " * space_size, r)

            if r.right != nil():
                print_tree_internal(r.right, space_size=space_size + 1)

        print(f"size = {self.size} / height = {self.height}")

        if self.root != nil():
            print_tree_internal(self.root)

    def __get_height(self):

        def get_height_internal(r):
            if r == nil():
                return -1

            height_of_left = get_height_internal(r.left)
            height_of_right = get_height_internal(r.right)

            return max(height_of_left, height_of_right) + 1

        if self.root == nil():
            return 0

        return get_height_internal(self.root)

    def __successor(self, z):
        if z.right:
            p = z.right
            while p.left:
                p = p.left

            return p
        else:
            p, q = z, z.parent
            while q and p == q.right:
                p = q
                q = q.parent

            return q

    def __transplant(self, p, q):
        if not p.parent:
            self.root = q
        elif p == p.parent.left:
            p.parent.left = q
        else:
            p.parent.right = q

        if q:
            q.parent = p.parent

    def __fix_after_insertion(self, z):
        while z.parent.color == Color.RED:
            if z.parent == z.parent.parent.left:
                y = z.parent.parent.right
                if y.color == Color.RED:
                    z.parent.color = Color.BLACK
                    z.parent.parent.color = Color.RED
                    y.color = Color.BLACK
                    z = z.parent.parent
                else:
                    if z == z.parent.right:
                        z = z.parent
                        self.__rotate_left(z)
                    z.parent.color = Color.BLACK
                    z.parent.parent.color = Color.RED
                    self.__rotate_right(z.parent.parent)
            else:
                y = z.parent.parent.left
                if y.color == Color.RED:
                    z.parent.color = Color.BLACK
                    y.color = Color.BLACK
                    z = z.parent.parent
                else:
                    if z == z.parent.left:
                        z = z.parent
                        self.__rotate_right(z)
                    z.parent.color = Color.BLACK
                    z.parent.parent.color = Color.RED
                    self.__rotate_left(z.parent.parent)

        self.root.color = Color.BLACK

    def __rotate_left(self, x):
        y = x.right

        x.right = y.left
        if y.left != nil():
            y.left.parent = x

        y.parent = x.parent
        if x.parent == nil():
            self.root = y
        elif x == x.parent.left:
            x.parent.left = y
        else:
            x.parent.right = y

        y.left = x
        x.parent = y

    def __rotate_right(self, x):
        y = x.left

        x.left = y.right
        if y.right != nil():
            y.right.parent = x

        y.parent = x.parent
        if x.parent == nil():
            self.root = y
        elif x == x.parent.left:
            x.parent.left = y
        else:
            x.parent.right = y

        y.right = x
        x.parent = y


if __name__ == "__main__":
    t = Tree()

    while True:
        print("1:put 2:remove 3:print > ", end="")
        if (op := int(input())) == 1:
            print("input key > ", end="")
            key = int(input())

            t.put(key)
        elif op == 2:
            print("input key > ", end="")
            key = int(input())

            t.remove(key)
        elif op == 3:
            t.print_tree()
        else:
            print("invalid operation")
