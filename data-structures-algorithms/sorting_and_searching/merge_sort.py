def merge(a: list[int], start, mid, end):
    left = a[start : mid + 1]
    right = a[mid + 1 : end + 1]
    i = j = 0
    k = start

    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            a[k] = left[i]
            i += 1
        else:
            a[k] = right[j]
            j += 1
        k += 1

    while i < len(left):
        a[k] = left[i]
        k += 1
        i += 1

    while j < len(right):
        a[k] = right[j]
        k += 1
        j += 1


def merge_sort(a: list[int], start, end):
    if start >= end:
        return
    mid = (start + end) // 2
    merge_sort(a, start, mid)
    merge_sort(a, mid + 1, end)
    merge(a, start, mid, end)


def main():
    a = [33, 2, 4, 5, 2, 1, 3, 44, 5, 33, 2]
    merge_sort(a, 0, len(a) - 1)
    print(a)


if __name__ == "__main__":
    main()
