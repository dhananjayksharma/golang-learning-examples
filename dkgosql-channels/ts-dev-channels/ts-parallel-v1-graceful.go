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

var maxCapacity = 2

func doParallelTask(ctx context.Context, targets []string) []error {
	sem := make(chan struct{}, maxCapacity)
	var mu sync.Mutex
	var errList []error
	var wg sync.WaitGroup

	for i, v := range targets {
		// Stop launching new goroutines if shutdown was signalled
		select {
		case <-ctx.Done():
			fmt.Printf("Shutdown: skipping remaining %d tasks\n", len(targets)-i)
			goto wait
		default:
		}

		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			// Re-check context — we may have been queued behind the semaphore
			// for a while; no point executing if shutdown started
			if ctx.Err() != nil {
				fmt.Printf("Shutdown: aborting task %d (%s)\n", i, v)
				return
			}

			fmt.Println("Task started:", v)
			err := execute(ctx, i, v)
			if err != nil {
				mu.Lock()
				errList = append(errList, err)
				mu.Unlock()
			}
		}(i, v)
	}

wait:
	fmt.Println("Waiting for in-flight tasks to finish...")
	wg.Wait()
	return errList
}

func execute(ctx context.Context, i int, task string) error {
	select {
	case <-time.After(2 * time.Second): // simulate work
		fmt.Println("Task done:", task)
		if i%2 == 0 {
			return fmt.Errorf("even index error at %d", i)
		}
		return fmt.Errorf("odd index error at %d", i)

	case <-ctx.Done():
		fmt.Println("Task cancelled mid-flight:", task)
		return fmt.Errorf("task %q cancelled: %w", task, ctx.Err())
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	fileList := []string{
		"file-8-apr-2026.json", "file-28-apr-2026.json", "file-20-apr-2026.json",
		"file-16-apr-2026.json", "file-12-apr-2026.json", "file-13-apr-2026.json",
		"file-14-apr-2026.json", "file-1-apr-2026.json", "file-11-apr-2026.json",
		"file-19-apr-2026.json",
	}

	errList := doParallelTask(ctx, fileList)

	fmt.Printf("\n--- Results ---\n")
	fmt.Printf("Errors (%d):\n", len(errList))
	for _, e := range errList {
		fmt.Println(" •", e)
	}
}
