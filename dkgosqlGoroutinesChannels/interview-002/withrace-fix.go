package main

import (
	"fmt"
	"sync"
)

func print(i int) {
	fmt.Println(i)
}

var cnt int = 0

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			cnt++
			current := cnt
			mu.Unlock()
			print(current)
		}(&mu)

	}
	wg.Wait()
}
