package main

import (
	"fmt"
	"sync"
)

func print(i int) {
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			print(i)
		}()

	}
	wg.Wait()
}
