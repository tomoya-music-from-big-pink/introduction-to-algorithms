class QueueEmptyException(Exception):
    pass


class Element:
    __slots__ = ["data", "next"]

    def __init__(self, data):
        self.data = data
        self.next = None


class Queue:
    __slots__ = ["size", "head", "tail"]

    def __init__(self):
        self.size = 0
        self.head = self.tail = None

    def enqueue(self, data):
        self.size += 1

        z = Element(data)

        if not self.head:
            self.head = z
        else:
            self.tail.next = z
        self.tail = z

    def dequeue(self):
        if not self.head:
            raise QueueEmptyException

        self.size -= 1

        dequeued = self.head
        result = dequeued.data
        self.head = dequeued.next

        return result

    def print_queue(self):
        print(f"size = {self.size} : ", end="")

        if not self.head:
            print("(empty)")

            return

        p = self.head
        while p:
            print(p.data, end=" -> " if p.next else "\n")

            p = p.next


if __name__ == "__main__":
    q = Queue()

    while True:
        print("1:enqueue 2:dequeue 3:print > ", end="")
        if (op := int(input())) == 1:
            print("input data > ", end="")
            data = int(input())

            q.enqueue(data)
        elif op == 2:
            try:
                dequeued = q.dequeue()
                print(f"dequeued data = {dequeued}")
            except QueueEmptyException:
                print("queue is empty")
        elif op == 3:
            q.print_queue()
        else:
            print("invalid operation")
