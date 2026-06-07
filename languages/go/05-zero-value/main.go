package main

import "fmt"

var (
	i int
	j bool
	k string
	l uint
	m rune
	n byte
	o uintptr
	p float64
	q complex64
	r *int
	s []int
)

func main() {
	fmt.Printf("i: %v\n", i)
	fmt.Printf("j: %v\n", j)
	fmt.Printf("k: %v\n", k)
	fmt.Printf("l: %v\n", l)
	fmt.Printf("m: %v\n", m)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("o: %v\n", o)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("q: %v\n", q)
	fmt.Printf("r: %v\n", r)
	fmt.Printf("s: %v\n", s)
}
