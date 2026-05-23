package main

import (
	"fmt"
	"sync"
)

func main() {
	values := []int{1, 3, 5, 6444, 4, 6, 7, 8888, 999}
	var wg sync.WaitGroup

	for _, v := range values {
		wg.Add(2)
		n := v
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(v)
		go func() {
			defer wg.Done()
			fmt.Printf("number %d\n", n)
		}()
	}

	wg.Wait()
}
