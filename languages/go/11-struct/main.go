package main

import "fmt"

type Fruit struct {
	name, color string
}

func (f Fruit) String() string {
	return fmt.Sprintf("name: %-10v color: %v", f.name, f.color)
}

func main() {
	fruit := Fruit{
		name:  "apple",
		color: "red",
	}
	fmt.Println(fruit)

	ptr := &fruit
	ptr.name = "cherry"

	fmt.Println(ptr)
}
