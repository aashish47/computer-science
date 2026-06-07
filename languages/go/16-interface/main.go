package main

import (
	"fmt"
	"math"
	"strings"
)

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return math.Pow(s.side, 2)
}

type Rectangle struct {
	length, breadth float64
}

func (r Rectangle) area() float64 {
	return r.breadth * r.length
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Triangle struct {
	height, base float64
}

func (t Triangle) area() float64 {
	return t.height * t.base * 0.5
}

func methods() {
	square := Square{
		side: 10,
	}

	rectangle := Rectangle{
		length:  2,
		breadth: 5,
	}

	circle := Circle{
		radius: 10,
	}

	triangle := Triangle{
		base:   10,
		height: 5,
	}

	shapes := map[string]Shape{
		"square":    square,
		"rectangle": rectangle,
		"circle":    circle,
		"triangle":  triangle,
	}

	divider := strings.Repeat("-", 50)

	// Print header
	fmt.Println(divider)
	fmt.Printf("%-10s | %-15s | %-10s | %s\n", "Shape", "Type", "Values", "Area")
	fmt.Println(divider)

	for name, shape := range shapes {
		typeStr := fmt.Sprintf("%T", shape)
		valueStr := fmt.Sprintf("%v", shape)
		fmt.Printf("%-10s | %-15s | %-10s | %.2f\n", name, typeStr, valueStr, shape.area())
	}
	fmt.Println(divider)
}

func nilInterface() {
	var i any
	fmt.Printf("%v %T", i, i)
}

func typeAssertion() {
	var i any = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f)
}

func typeSwitches(i any) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case int:
		fmt.Printf("%v time 10 is %v\n", v, v*10)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	// methods()
	// nilInterface()
	// typeAssertion()
	typeSwitches(10)
	typeSwitches("hello")
	typeSwitches(10.0)
}
