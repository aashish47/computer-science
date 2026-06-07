package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	switch time.Sunday {
	case today:
		fmt.Println("today is sunday")
	case (today + 1) % 7:
		fmt.Println("tomorrow is sunday")
	case (today + 2) % 7:
		fmt.Println("day after tomorrow is sunday")
	default:
		fmt.Println("sunday is too far")
	}

}
