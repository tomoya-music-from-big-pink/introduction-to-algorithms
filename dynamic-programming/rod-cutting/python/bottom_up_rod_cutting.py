import sys


def cut_rod(prices, length):
    results = [-sys.maxsize for _ in range(length + 1)]
    slices = [0 for _ in range(length + 1)]

    results[0] = 0

    for j in range(1, length):
        result = -sys.maxsize
        for i in range(1, j + 1):
            if result < prices[i] + results[j - i]:
                result = prices[i] + results[j - i]
                slices[j] = i

        results[j] = result

    return results, slices


def print_optimal_strategy(slices, length):
    print("(", end="")

    rest = length
    while rest > 0:
        print(slices[rest], end="")

        if rest - slices[rest] > 0:
            print(" + ", end="")

        rest -= slices[rest]

    print(")", end="")


if __name__ == "__main__":
    prices = [0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30]

    print("  i |", end="")
    for i in range(len(prices)):
        print(f"{i:3}|", end="")
    print("")
    print("p[i]|", end="")
    for p in prices:
        print(f"{p:3}|", end="")
    print("")

    results, slices = cut_rod(prices, len(prices))
    for n in range(len(prices)):
        print(f"length = {n} / result = {results[n]} ", end=" ")

        if n != 0:
            print_optimal_strategy(slices, n)

        print("")
