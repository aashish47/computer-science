package main

import "fmt"

func index[T comparable](list []T, target T) int {
	for i, val := range list {
		if val == target {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	val  T
	next *List[T]
}

func main() {
	fruits := []string{"apples", "oranges"}
	fmt.Printf("%v\n", index(fruits, "oranges"))
	fmt.Printf("%v\n", index(fruits, "mangoes"))

	nums := []int{10, 20}
	fmt.Printf("%v\n", index(nums, 10))
	fmt.Printf("%v\n", index(nums, 30))
}
