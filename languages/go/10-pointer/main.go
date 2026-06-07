package main

import "fmt"

func main() {
	i := 10
	j := &i
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: address: %v content: %v", j, *j)
}
