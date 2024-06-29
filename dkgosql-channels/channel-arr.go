package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var numcpu = runtime.NumCPU()

func _init() {
	runtime.GOMAXPROCS(2) // Try to use all available CPUs.
}
func main() {
	runtime.GOMAXPROCS(1)
	start := time.Now()
	// Code to measure

	num := 10000
	var wg sync.WaitGroup
	n := 2

	for i := 0; i < n; i++ {
		wg.Add(n)
		go worker(&wg, num)
		go worker(&wg, num)
		// go worker(&wg, num)
		// go worker(&wg, num)
		// go worker(&wg, num)
		// go worker(&wg, num)
		// go worker(&wg, num)
		// go worker(&wg, num)
		fmt.Printf("\ngo routine:%d \n", runtime.NumGoroutine())

	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("\nAll done:%s, elapsed time:%v, total of CPU:%d, Used CPU:%d go routine:%d \n", "true", duration, numcpu, runtime.NumCgoCall(), runtime.NumGoroutine())
}

func worker(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		_ = i
		fmt.Printf("D:%d", i)
	}
}
