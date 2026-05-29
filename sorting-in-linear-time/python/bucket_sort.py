import random


def sort(data):
    table = [[] for _ in range(max(data) // 10 + 1)]

    for d in data:
        table[d // 10].append(d)

    print("--- table before sorting ---")
    for i, t in enumerate(table):
        print(i, ":", t)

    for t in table:
        t = t.sort()

    print("--- table after sorting ---")
    for i, t in enumerate(table):
        print(i, ":", t)

    data.clear()
    for t in table:
        data.extend(t)


if __name__ == "__main__":
    data = [random.randint(1, 100) for _ in range(100)]

    print("--- before ---")
    print(data)

    print("--- sort ---")
    sort(data)

    print("--- after ---")
    print(data)
