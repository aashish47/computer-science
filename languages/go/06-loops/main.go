package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Printf("%-2v", i)
	}
	fmt.Println()
	i := 10
	for i < 20 {
		fmt.Printf("%v ", i)
		i += 2
	}
}
