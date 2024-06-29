package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a channel for communication between the main thread and the goroutine
	ch := make(chan int)

	// Use a wait group to synchronize the termination of the goroutine
	var wg sync.WaitGroup
	wg.Add(1)

	// Goroutine to print numbers repeatedly
	go func() {
		defer wg.Done()
		for i := 1; ; i++ {
			fmt.Println(i)
			ch <- i // Send the number to the channel
		}
	}()

	// Main thread loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("Main Thread: %d\n", i)
		num := <-ch // Receive the number from the channel
		if num == 10 {
			break // Exit the loop when the goroutine reaches 10
		}
	}

	close(ch) // Close the channel to signal the goroutine to stop
	wg.Wait() // Wait for the goroutine to finish
}
