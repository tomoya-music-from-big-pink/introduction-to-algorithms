import sys


def matrix_chain_order(p, start_idx, end_idx):
    results = [[sys.maxsize for _ in range(len(p) + 1)] for _ in range(len(p) + 1)]

    return matrix_chain_order_internal(p, start_idx, end_idx, results)


def matrix_chain_order_internal(p, start_idx, end_idx, results):
    if results[start_idx][end_idx] < sys.maxsize:
        return results[start_idx][end_idx]

    if start_idx == end_idx:
        result = 0
    else:
        result = sys.maxsize
        for k in range(start_idx, end_idx):
            result = min(
                result,
                matrix_chain_order_internal(p, start_idx, k, results)
                + matrix_chain_order_internal(p, k + 1, end_idx, results)
                + p[start_idx] * p[k + 1] * p[end_idx + 1],
            )

    results[start_idx][end_idx] = result

    return result


if __name__ == "__main__":
    p = [30, 35, 15, 5, 10, 20, 25]

    print("  i |", end="")
    for i in range(len(p) - 1):
        print(f"{i:^7}|", end="")
    print("")
    print("A[i]|", end="")
    for i in range(len(p) - 1):
        print(f"{p[i]:3}*{p[i + 1]:3}|", end="")
    print("")

    diff = 0
    while diff < len(p) - 1:
        print(f"--- diff = {diff} ---")

        i, j = 0, diff
        while j < len(p) - 1:
            result = matrix_chain_order(p, i, j)

            print(f"i = {i}, j = {j} / result = {result}")

            i += 1
            j += 1

        diff += 1
