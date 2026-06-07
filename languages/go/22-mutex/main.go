package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	m  map[string]int
}

func (c *Counter) inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key]++
}

func (c *Counter) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.m[key]
}

func main() {
	counter := Counter{
		m: make(map[string]int),
	}

	for range 1000 {
		go counter.inc("someKey")

	}
	time.Sleep(time.Second)
	fmt.Println(counter.value("someKey"))
}
