package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %v working on job: %v\n", i, job)
		time.Sleep(2 * time.Second)
		results <- job
	}

}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := range 5 {
		wg.Add(1)
		go worker(i, &wg, jobs, results)
	}

	for i := range 10 {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("job %v completed\n", result)
	}
}
