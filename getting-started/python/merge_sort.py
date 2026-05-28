import random


def sort(data):
    merge_sort(data, 0, len(data) - 1)


def merge_sort(data, start_idx, end_index):
    if start_idx < end_index:
        middle_idx = (start_idx + end_index) // 2

        merge_sort(data, start_idx, middle_idx)
        merge_sort(data, middle_idx + 1, end_index)

        merge(data, start_idx, middle_idx, end_index)


def merge(data, start_idx, middle_idx, end_index):
    first_half = data[start_idx : middle_idx + 1]
    second_half = data[middle_idx + 1 : end_index + 1]

    i = j = 0
    for k in range(start_idx, end_index + 1):
        if i < len(first_half) and j < len(second_half):
            if first_half[i] < second_half[j]:
                data[k] = first_half[i]
                i += 1
            else:
                data[k] = second_half[j]
                j += 1
        elif i < len(first_half):
            data[k] = first_half[i]
            i += 1
        else:
            data[k] = second_half[j]
            j += 1

    print(data)


if __name__ == "__main__":
    data = [random.randint(1, 20) for _ in range(20)]

    print("--- before ---")
    print(data)

    print("--- sort ---")
    sort(data)

    print("--- after ---")
    print(data)
