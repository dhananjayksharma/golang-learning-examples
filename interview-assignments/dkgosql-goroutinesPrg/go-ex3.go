package main

import (
	"fmt"
	"sync"
	"time"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner\n")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second*2)
	fmt.Print("\nI am second runner\n")
}

func execute() {
	
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go runner1(wg)
	go runner2(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func main() {
	start := time.Now()
	// Launching both the runners
	execute()
	// Code to measure
	duration := time.Since(start)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)
}
