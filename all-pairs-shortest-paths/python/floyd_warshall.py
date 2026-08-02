import copy
import sys


def floyd_warshall(W):
    n = len(W)

    D = copy.deepcopy(W)
    for k in range(n):
        for i in range(n):
            for j in range(n):
                D[i][j] = min(D[i][j], D[i][k] + D[k][j])

    return D

def print_matrix(A):
    if not A:
        return

    for row in A:
        for i in range(len(row)):
            print(f"{row[i]:3d}" if row[i] < sys.maxsize else "INF", end=", ")
        print("")

if __name__ == '__main__':
    W = [
        [0, 3, 8, sys.maxsize, -4],
        [sys.maxsize, 0, sys.maxsize, 1, 7],
        [sys.maxsize, 4, 0, sys.maxsize, sys.maxsize],
        [2, sys.maxsize, -5, 0, sys.maxsize],
        [sys.maxsize, sys.maxsize, sys.maxsize, 6, 0],
    ]

    print("--- graph ---")
    print_matrix(W)

    print("--- Floyd-Warshall algorithm ---")
    D = floyd_warshall(W)

    print("--- result ---")
    print_matrix(D)
