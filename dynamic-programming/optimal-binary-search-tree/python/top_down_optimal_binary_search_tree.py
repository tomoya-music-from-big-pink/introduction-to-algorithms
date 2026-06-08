import sys


def optimal_bst(p, q, i, j):
    results = [[sys.maxsize for _ in range(
        max(len(p), len(q) + 1))] for _ in range(max(len(p), len(q) + 1))]

    return optimal_bst_internal(p, q, i, j, results)


def optimal_bst_internal(p, q, i, j, results):
    if results[i][j] < sys.maxsize:
        return results[i][j]

    if j == i - 1:
        result = q[i - 1]
    else:
        w = q[i - 1]
        for l in range(i, j + 1):
            w = w + p[l] + q[l]
        w = round(w, 2)

        result = sys.maxsize
        for r in range(i, j + 1):
            result = min(result, optimal_bst_internal(
                p, q, i, r - 1, results) + optimal_bst_internal(p, q, r + 1, j, results) + w)

    results[i][j] = result

    return result


if __name__ == '__main__':
    p = [0.0, 0.15, 0.10, 0.05, 0.10, 0.20]
    q = [0.05, 0.10, 0.05, 0.05, 0.05, 0.10]

    print("  i |", end='')
    for i in range(len(p)):
        print(f"{i:^5}|", end='')
    print("")
    print("p[i]|", end='')
    for _p in p:
        print(f"{_p:>5}|", end='')
    print("")
    print("q[i]|", end='')
    for _q in q:
        print(f"{_q:>5}|", end='')
    print("")

    diff = -1
    while diff < len(p) - 1:
        print(f"--- diff = {diff} ---")

        i = 1
        j = i + diff
        while j < len(p):
            result = optimal_bst(p, q, i, j)

            print(f"i = {i}, j = {j} / result = {result:.2f}")

            i += 1
            j += 1

        diff += 1
