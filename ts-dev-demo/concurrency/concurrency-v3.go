package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func workerpool(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum += i
	}

	fmt.Println("Worker:", id, "done", sum)
}

func main() {
	start := time.Now() // start time
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup

	for i := 1; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go workerpool(i, &wg)
	}

	wg.Wait()

	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	duration := time.Since(start) // end - start

	fmt.Println("Execution time:", duration)
}
