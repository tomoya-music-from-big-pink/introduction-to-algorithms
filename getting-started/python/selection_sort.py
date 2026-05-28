import random


def sort(data):
    for i in range(len(data) - 1):
        idx_of_min = i

        for j in range(i + 1, len(data)):
            if data[j] < data[idx_of_min]:
                idx_of_min = j

        if idx_of_min != i:
            data[i], data[idx_of_min] = data[idx_of_min], data[i]

        print(data)


if __name__ == "__main__":
    data = [random.randint(1, 20) for _ in range(20)]

    print("--- before ---")
    print(data)

    print("--- sort ---")
    sort(data)

    print("--- after ---")
    print(data)
