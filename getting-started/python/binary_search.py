def search(data, key):
    start_idx = 0
    end_idx = len(data) - 1

    while start_idx <= end_idx:
        middle_idx = (start_idx + end_idx) // 2

        if data[middle_idx] == key:
            return middle_idx
        elif data[middle_idx] < key:
            start_idx = middle_idx + 1
        else:
            end_idx = middle_idx - 1

    return -1


if __name__ == "__main__":
    data = [i**2 for i in range(20)]

    print("--- data ---")
    print(data)

    print("input key > ", end="")
    key = int(input())

    result = search(data, key)
    if result != -1:
        print(f"result = {result}")
    else:
        print(f"{key} is not found.")
