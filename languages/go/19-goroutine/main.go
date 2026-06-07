package main

import (
	"fmt"
	"time"
)

func greet(s string) {
	for range 5 {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	start := time.Now()
	go greet("Hello")
	greet("World")
	elapsed := time.Since(start)
	fmt.Printf("total time: %v", elapsed)
}
