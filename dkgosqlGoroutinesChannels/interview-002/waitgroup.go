package main

import (
	"fmt"
	"sync"
)

func print(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go print(&wg, i)
	}
	wg.Wait()
}
