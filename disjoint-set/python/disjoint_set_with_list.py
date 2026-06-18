class Element:
    __slots__ = ["data", "next", "prev"]

    def __init__(self, data):
        self.data = data
        self.next = self.prev = None


class LinkedList:
    __slots__ = ["head", "tail"]

    def __init__(self):
        self.head = self.tail = None


def make_set(data):
    ll = LinkedList()
    ll.head = ll.tail = Element(data)

    return ll


def union(ch1, ch2, sets):
    parent_of_x = find_set(ch1, sets)
    parent_of_y = find_set(ch2, sets)

    if parent_of_x == parent_of_y:
        return

    x = sets[parent_of_x]
    y = sets[parent_of_y]

    if parent_of_x < parent_of_y:
        x.tail.next = y.head
        x.tail = y.tail
        del sets[ch2]


def find_set(data, sets):
    for v in sets.values():
        found = False
        p = v.head
        while p:
            if p.data == data:
                found = True

                break

            p = p.next

        if found:
            return v.head.data

    return None


def print_sets(sets):
    for k, v in sets.items():
        print(k, ":", end=" ")
        p = v.head
        while p:
            print(p.data, end=" -> " if p.next else "")

            p = p.next

        print("")


if __name__ == "__main__":
    characters = "abcdefghij"

    sets = {}
    for ch in characters:
        sets[ch] = make_set(ch)

    print("--- after initialize ---")
    print_sets(sets)

    print("union b and d")
    union("b", "d", sets)
    print_sets(sets)
    print("union e and g")
    union("e", "g", sets)
    print_sets(sets)
    print("union a and c")
    union("a", "c", sets)
    print_sets(sets)
    print("union h and i")
    union("h", "i", sets)
    print_sets(sets)
    print("union a and b")
    union("a", "b", sets)
    print_sets(sets)
    print("union e and f")
    union("e", "f", sets)
    print_sets(sets)
    print("union b and c")
    union("b", "c", sets)
    print_sets(sets)
