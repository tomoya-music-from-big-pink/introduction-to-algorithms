import sys


def common_lcs(x, y, i, j):
    results = [
        [-sys.maxsize for _ in range(max(len(x), len(y)) + 1)]
        for _ in range(max(len(x), len(y)) + 1)
    ]

    return common_lcs_internal(x, y, i, j, results)


def common_lcs_internal(x, y, i, j, results):
    if results[i][j] > -sys.maxsize:
        return results[i][j]

    if i == 0 and j == 0:
        result = 1 if x[i] == y[j] else 0
    elif i == 0:
        result = 1 if x[i] == y[j] else common_lcs_internal(x, y, i, j - 1, results)
    elif j == 0:
        result = 1 if x[i] == y[j] else common_lcs_internal(x, y, i - 1, j, results)
    else:
        if x[i] == y[j]:
            result = common_lcs_internal(x, y, i - 1, j - 1, results) + 1
        else:
            result = max(
                common_lcs_internal(x, y, i - 1, j, results),
                common_lcs_internal(x, y, i, j - 1, results),
            )

    results[i][j] = result

    return result


if __name__ == "__main__":
    x = "ABCBDAB"
    y = "BDCABA"

    for i in range(len(x)):
        print(f"--- i = {i} ---")
        for j in range(len(y)):
            result = common_lcs(x, y, i, j)

            print(
                f"i = {i}, j = {j}, x = {x[: i + 1]}, y = {y[: j + 1]} / result = {result}"
            )
