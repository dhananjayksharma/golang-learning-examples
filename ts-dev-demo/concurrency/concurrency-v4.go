package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func cpuTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	total := 0
	for i := 0; i < 500000000; i++ {
		total += i
	}

	fmt.Println("Task done:", id, total)
}

func runTest(maxProcs int) {
	runtime.GOMAXPROCS(maxProcs)

	start := time.Now()

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go cpuTask(i, &wg)
	}

	wg.Wait()

	fmt.Printf("GOMAXPROCS=%d, Time=%v\n", maxProcs, time.Since(start))
}

func main() {
	fmt.Println("NumCPU:", runtime.NumCPU())

	runTest(1)
	runTest(runtime.NumCPU())
}
