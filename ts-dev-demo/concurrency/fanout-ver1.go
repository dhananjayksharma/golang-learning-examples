package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker-%d processing job %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup
	numberOfWorker := 3
	wg.Add(numberOfWorker)

	for w := 1; w <= numberOfWorker; w++ {
		go worker(w, jobs, &wg)
	}
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
