package main

import (
	"fmt"
	"sync"
)

// BAD: concurrent writes to counter — data race
var counter int

func increment() {

	counter++ // NOT atomic: read-modify-write, 3 ops
}

// Run with: go run -race main.go
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println(counter) // not reliably 1000
}
