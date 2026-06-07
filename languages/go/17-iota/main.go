package main

import (
	"fmt"
)

const (
	Bit0 = 1 << iota
	Bit1
	Bit2
	Bit3
)

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
)

func main() {
	fmt.Println(Bit0)
	fmt.Println(Bit1)
	fmt.Println(Bit2)
	fmt.Println(Bit3)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
