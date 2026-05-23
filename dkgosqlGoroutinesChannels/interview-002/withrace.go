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
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt++
			print(cnt)
		}()

	}
	wg.Wait()
}
