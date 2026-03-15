package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, jobs <-chan string, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[worker-%d] stop-fast exit\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("[worker-%d] jobs closed\n", id)
				return
			}
			fmt.Printf("[worker-%d] processing %s\n", id, job)
			time.Sleep(200 * time.Millisecond) // simulate work
		}
	}
}

func waitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan string, 100)
	var wg sync.WaitGroup

	numWorkers := 5
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, jobs, &wg, i)
	}

	// producer
	go func() {
		defer close(jobs)
		for i := 1; i <= 510; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("[producer] cancel, stop producing")
				return
			case jobs <- fmt.Sprintf("job-%d", i):
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("[main] shutdown signal received")

	ok := waitWithTimeout(&wg, 2*time.Millisecond)
	if !ok {
		fmt.Println("[main] timeout waiting workers; exiting anyway")
		return
	}
	fmt.Println("[main] graceful shutdown complete")
}
