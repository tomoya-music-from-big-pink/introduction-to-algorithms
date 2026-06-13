import heapq


class Element:
    __slots__ = ["ch", "freq", "left", "right", "parent"]

    def __init__(self, ch, freq):
        self.ch = ch
        self.freq = freq
        self.left = self.right = self.right = None

    def __lt__(self, other):
        return self.freq < other.freq


def print_tree(root):

    def print_tree_internal(r, space_size=0):
        if not r:
            return

        print_tree_internal(r.left, space_size=space_size + 1)
        print("  " * space_size, f"{r.ch}({r.freq})")
        print_tree_internal(r.right, space_size=space_size + 1)

    print_tree_internal(root)


def huffman(characters, frequencies):
    elements = []
    for ch, freq in zip(characters, frequencies):
        elements.append(Element(ch, freq))

    heapq.heapify(elements)

    while len(elements) > 1:
        z = Element("", 0)
        x = heapq.heappop(elements)
        y = heapq.heappop(elements)

        z.freq = x.freq + y.freq
        z.left = x
        z.right = y
        x.parent = y.parent = z

        heapq.heappush(elements, z)

    return heapq.heappop(elements)


if __name__ == "__main__":
    characters = "abcdef"
    frequencies = [45, 13, 12, 16, 9, 5]

    print(" ch |", end="")
    for ch in characters:
        print(f"{ch:^3}|", end="")
    print("")
    print("freq|", end="")
    for freq in frequencies:
        print(f"{freq:^3}|", end="")
    print("")

    root = huffman(characters, frequencies)
    print_tree(root)
