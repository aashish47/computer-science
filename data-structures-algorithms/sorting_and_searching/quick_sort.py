def quick_sort(a: list[int], left, right):
    if left < right:
        index = partition(a, left, right)
        quick_sort(a, left, index - 1)
        quick_sort(a, index + 1, right)


def partition(a: list[int], left, right):
    small = left
    for i in range(left, right):
        if a[i] <= a[right]:
            a[i], a[small] = a[small], a[i]
            small += 1
    a[small], a[right] = a[right], a[small]
    return small


# def quick_sort(a: list[int]):
#     if len(a) <= 1:
#         return a
#     pivot = a[len(a) // 2]
#     less_than_pivot = [x for x in a if x < pivot]
#     equal_to_pivot = [x for x in a if x == pivot]
#     greater_than_pivot = [x for x in a if x > pivot]

#     return quick_sort(less_than_pivot) + equal_to_pivot + quick_sort(greater_than_pivot)


def main():
    a = [33, 2, 4, 5, 2, 1, 3, 44, 5, 33, 2]
    quick_sort(a, 0, len(a) - 1)
    print(a)


if __name__ == "__main__":
    main()
