package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(ctx context.Context, jobs <-chan string, wg *sync.WaitGroup, id string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] shutdown signal received\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("[%s] jobs channel closed\n", id)
				return
			}
			fmt.Printf("[%s] processing: %s\n", id, job)
		}
	}
}

func main() {
	fmt.Println("starting now")
	runtime.GOMAXPROCS(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan string)
	var wg sync.WaitGroup

	numWorkers := 5
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, jobs, &wg, fmt.Sprintf("worker-%d", i))
	}

	// simulate shutdown happening early
	go func() {
		time.Sleep(15 * time.Minute)
		cancel()
	}()

	jobAllocationCnt := 0
	// producer
	for i := 1; i <= 5000000; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("producer stopping early")
			close(jobs)
			wg.Wait()
			fmt.Println("ending... now")
			return
		case jobs <- fmt.Sprintf("job id: %d", i):
			fmt.Printf("Job assigned: %v\n", i)
			jobAllocationCnt++
		}
	}

	close(jobs)
	wg.Wait()
	fmt.Printf("Job Allocation done Count: %d\n", jobAllocationCnt)
	fmt.Println("ending... now")
	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("NumCPU:", runtime.NumCPU())

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	fmt.Println("NumThread:", runtime.NumCgoCall()) // if using cgo

}
