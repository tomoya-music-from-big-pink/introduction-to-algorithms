class Element:
    __slots__ = ["data", "next", "prev"]

    def __init__(self, data):
        self.data = data
        self.next = self.prev = None


class LinkedList:
    __slots__ = ["size", "head", "tail"]

    def __init__(self):
        self.size = 0
        self.head = self.tail = None

    def add(self, data):
        self.size += 1

        z = Element(data)

        if not self.head:
            self.head = z
        else:
            self.tail.next = z
        z.prev = self.tail
        self.tail = z

    def remove(self, data):
        z = self.__contains(data)
        if not z:
            print(f"{data} is not found in the list")

            return

        self.size -= 1

        if z.next:
            z.next.prev = z.prev
        else:
            self.tail = z.prev
        if z.prev:
            z.prev.next = z.next
        else:
            self.head = z.next

    def __contains(self, data):
        p = self.head
        while p and p.data != data:
            p = p.next

        return p

    def print_list(self):
        print(f"size = {self.size}", end="")

        if not self.head:
            print(" (empty)")

            return

        print("\n--- from head to tail ---")
        p = self.head
        while p:
            print(p.data, end=" -> " if p.next else "\n")

            p = p.next

        print("--- from tail to head ---")
        p = self.tail
        while p:
            print(p.data, end=" <- " if p.prev else "\n")

            p = p.prev


if __name__ == "__main__":
    ll = LinkedList()

    while True:
        print("1:add 2:remove 3:print > ", end="")
        if (op := int(input())) == 1:
            print("input data > ", end="")
            data = int(input())

            ll.add(data)
        elif op == 2:
            print("input data > ", end="")
            data = int(input())

            ll.remove(data)
        elif op == 3:
            ll.print_list()
        else:
            print("invalid operation")
