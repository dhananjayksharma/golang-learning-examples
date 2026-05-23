package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped:", ctx.Err())
				return
			default:
				fmt.Println("goroutine is running")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(300 * time.Millisecond)
}
