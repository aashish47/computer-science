package main

import (
	"fmt"
	"time"
)

func tick() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	for {
		select {
		case <-tick:
			fmt.Printf("[%6v] Tick\n", elapsed())
		case <-boom:
			fmt.Printf("[%6v] Boom\n", elapsed())
			return
		default:
			fmt.Printf("[%6v] ...\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func fibonacci() {
	nums := make(chan int)
	quit := make(chan int)
	go func() {
		for i := range 10 {
			fmt.Printf("%v: %v\n", i+1, <-nums)
		}
		quit <- 0
	}()

	x, y := 0, 1
	for {
		select {
		case nums <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Print("quit")
			return
		}
	}
}

func main() {
	// fibonacci()
	tick()

}
