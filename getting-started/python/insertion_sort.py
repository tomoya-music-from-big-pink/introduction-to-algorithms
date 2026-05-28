import random


def sort(data):
    for j in range(1, len(data)):
        key = data[j]
        i = j - 1
        while i >= 0 and data[i] > key:
            data[i], data[i + 1] = data[i + 1], data[i]

            i -= 1

        data[i + 1] = key

        print(data)


if __name__ == "__main__":
    data = [random.randint(1, 20) for _ in range(20)]

    print("--- before ---")
    print(data)

    print("--- sort ---")
    sort(data)

    print("--- after ---")
    print(data)
