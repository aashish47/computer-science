package main

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	adder5 := adder(5)
	println(adder5(10))
	println(adder5(20))
}
