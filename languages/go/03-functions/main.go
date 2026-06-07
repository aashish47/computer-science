package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func expand(a int) (x, y int) {
	x, y = a+10, a-10
	return
}

func main() {
	a, b := 5, 2
	fmt.Printf("a: %v, b: %v\n", a, b)
	fmt.Printf("sum: %v\n", add(a, b))
	a, b = swap(a, b)
	fmt.Printf("swap a: %v, b: %v\n", a, b)
	x, y := expand(a)
	fmt.Printf("expand a: %v %v", x, y)
}
