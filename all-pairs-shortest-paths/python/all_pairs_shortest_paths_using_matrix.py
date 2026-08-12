import copy
import sys


def extended_shortest_paths(L, W):
    new_L = copy.deepcopy(L)
    n = len(W)

    for i in range(n):
        for j in range(n):
            l = sys.maxsize
            for k in range(n):
                l = min(l, L[i][k] + W[k][j])

            new_L[i][j] = l

    return new_L


def all_pairs_shortest_paths(W):
    L = copy.deepcopy(W)
    n = len(W)

    for i in range(n):
        L = extended_shortest_paths(L, W)

        print(f"--- i = {i} ---")
        print_matrix(L)

    return L


def print_matrix(A):
    if not A:
        return

    for row in A:
        for i in range(len(row)):
            print(f"{row[i]:3d}" if row[i] < sys.maxsize else "INF", end=", ")
        print("")


if __name__ == "__main__":
    W = [
        [0, 3, 8, sys.maxsize, -4],
        [sys.maxsize, 0, sys.maxsize, 1, 7],
        [sys.maxsize, 4, 0, sys.maxsize, sys.maxsize],
        [2, sys.maxsize, -5, 0, sys.maxsize],
        [sys.maxsize, sys.maxsize, sys.maxsize, 6, 0],
    ]

    print("--- graph ---")
    print_matrix(W)

    L = all_pairs_shortest_paths(W)

    print("--- result ---")
    print_matrix(L)
