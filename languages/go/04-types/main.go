package main

import "fmt"

func main() {
	isTrue := false
	name := "aashish"
	amount := 100
	ptr := &amount
	price := 20.0
	fmt.Printf("%-10s-> type: %-10T value: %v\n", "isTrue", isTrue, isTrue)
	fmt.Printf("%-10s-> type: %-10T value: %v\n", "name", name, name)
	fmt.Printf("%-10s-> type: %-10T value: %v\n", "amount", amount, amount)
	fmt.Printf("%-10s-> type: %-10T value: %v\n", "ptr", ptr, ptr)
	fmt.Printf("%-10s-> type: %-10T value: %.2f\n", "price", price, price)
	price = float64(amount)
	fmt.Printf("%-10s-> type: %-10T value: %.2f\n", "price", price, price)
}
