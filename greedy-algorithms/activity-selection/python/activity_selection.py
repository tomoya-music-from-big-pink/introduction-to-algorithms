def select_activities(start, finish):
    activities = []

    activities.append(0)

    i = 1
    finished = finish[0]
    while i < len(finish):
        if start[i] > finished:
            activities.append(i)

            finished = finish[i]

        i += 1

    return activities


if __name__ == "__main__":
    start = [1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12]
    finish = [4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16]

    print("  i |", end="")
    for i in range(len(start)):
        print(f"{i:>3}|", end="")
    print("")
    print("s[i]|", end="")
    for s in start:
        print(f"{s:>3}|", end="")
    print("")
    print("f[i]|", end="")
    for f in finish:
        print(f"{f:>3}|", end="")
    print("")

    activities = select_activities(start, finish)
    print(f"result = {', '.join(map(str, activities))}")
