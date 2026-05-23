package main

import (
	"fmt"
	"runtime"
	"sync"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println("Worker:", id, "Step:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}

	wg.Wait()

	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
