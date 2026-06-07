package main

import "fmt"

func print(list []int) {
	fmt.Printf("slice: %v length: %v capacity: %v\n", list, len(list), cap(list))
}

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("array: %v\n", nums)

	slice := nums[1:3]
	print(slice)
	slice[0] = 0
	print(slice)

	a := make([]int, 5, 10)
	print(a)

	a = append(a, 10)
	print(a)

	b := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, list := range b {
		for _, val := range list {
			fmt.Printf("%v ", val)
		}
		fmt.Println()
	}
}
