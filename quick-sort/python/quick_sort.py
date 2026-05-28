import random


def sort(data):
    quick_sort(data, 0, len(data) - 1)


def quick_sort(data, start_idx, end_idx):
    if start_idx < end_idx:
        partition_idx = partition(data, start_idx, end_idx)

        quick_sort(data, start_idx, partition_idx - 1)
        quick_sort(data, partition_idx + 1, end_idx)


def partition(data, start_idx, end_idx):
    pivot = data[end_idx]

    i = start_idx
    for j in range(start_idx, end_idx):
        if data[j] < pivot:
            data[i], data[j] = data[j], data[i]

            i += 1

    data[i], data[end_idx] = data[end_idx], data[i]

    print(data)

    return i


if __name__ == "__main__":
    data = [random.randint(1, 20) for _ in range(20)]

    print("--- before ---")
    print(data)

    print("--- sort ---")
    sort(data)

    print("--- after ---")
    print(data)
