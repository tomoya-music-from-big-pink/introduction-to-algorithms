class Element:
    __slots__ = ["key", "next", "prev"]

    def __init__(self, key):
        self.key = key
        self.next = self.prev = None


class Hashtable:
    __slots__ = ["bucket", "bucket_size"]

    def __init__(self, bucket_size=13):
        self.bucket = [None for _ in range(bucket_size)]
        self.bucket_size = bucket_size

    def put(self, key):
        hash_val = self.__calculate_hash_value(key)

        z = Element(key)
        if self.bucket[hash_val] is not None:
            self.bucket[hash_val].prev = z
        z.next = self.bucket[hash_val]
        self.bucket[hash_val] = z

    def remove(self, key):
        z = self.__contains(key)
        if z is None:
            print(f"{key} is not found in the hash table")

            return

        hash_val = self.__calculate_hash_value(key)
        if z.next:
            z.next.prev = z.prev
        if z.prev:
            z.prev.next = z.next
        else:
            self.bucket[hash_val] = z.next

    def __contains(self, key):
        hash_val = self.__calculate_hash_value(key)

        p = self.bucket[hash_val]
        while p and p.key != key:
            p = p.next

        return p

    def print_hash_table(self):
        for i in range(self.bucket_size):
            print(i, ": ", end="")

            p = self.bucket[i]
            while p:
                print(p.key, end=" -> " if p.next else "\n")

                p = p.next

    def __calculate_hash_value(self, key):
        return key % self.bucket_size


if __name__ == "__main__":
    ht = Hashtable()
    for i in range(100):
        ht.put(i)
    ht.print_hash_table()

    while True:
        print("1:put 2:remove 3:print > ", end="")
        if (op := int(input())) == 1:
            print("input key > ", end="")
            key = int(input())

            ht.put(key)
        elif op == 2:
            print("input key > ", end="")
            key = int(input())

            ht.remove(key)
        elif op == 3:
            ht.print_hash_table()
        else:
            print("invalid operation")
