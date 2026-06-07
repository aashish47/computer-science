package main

import "fmt"

type Department struct {
	faculty, students int
}

func main() {
	m := make(map[string]Department)
	m["physics"] = Department{
		10, 100,
	}
	fmt.Println(m["physics"])

	n := map[string]Department{
		"physics": {10, 100},
		"maths":   {20, 50},
	}

	for key, value := range n {
		fmt.Printf("key: %-10v value: %v\n", key, value)
	}

	v, ok := n["science"]
	fmt.Printf("value: %v ok? %v\n", v, ok)

	delete(n, "physics")
	fmt.Println(n)
}
