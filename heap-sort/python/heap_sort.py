import random


def sort(data):
    build_max_heap(data)

    print(data)

    for i in reversed(range(1, len(data))):
        data[0], data[i] = data[i], data[0]

        heapify(data, 0, i - 1)

        print(data)


def build_max_heap(data):
    for i in reversed(range(len(data) // 2)):
        heapify(data, i, len(data) - 1)


def heapify(data, target_idx, heap_idx):
    largest_idx = target_idx

    left_idx = 2 * target_idx + 1
    if left_idx <= heap_idx and data[left_idx] > data[largest_idx]:
        largest_idx = left_idx

    right_idx = 2 * target_idx + 2
    if right_idx <= heap_idx and data[right_idx] > data[largest_idx]:
        largest_idx = right_idx

    if largest_idx != target_idx:
        data[target_idx], data[largest_idx] = data[largest_idx], data[target_idx]

        heapify(data, largest_idx, heap_idx)


if __name__ == "__main__":
    data = [random.randint(1, 20) for _ in range(20)]

    print("--- before ---")
    print(data)

    print("--- sort ---")
    sort(data)

    print("--- after ---")
    print(data)
