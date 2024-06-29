package main

import (
	"fmt"
	"sync"
)

var data int

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		defer wg.Done()
		mu.Lock()
		data++
		fmt.Printf("Step 1 the value is %v.\n", data)
		mu.Unlock()

	}(&wg, &mu)

	wg.Wait()

	if data == 0 {
		fmt.Printf("Step 2 the value is %v.\n", data)
	}

	fmt.Printf("Step 3 the value is %v.\n", data)

}

// 1
// 2
// 3
