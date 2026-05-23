package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		go func(id int) {
			time.Sleep(2 * time.Second)
			fmt.Println("done:", id)
		}(i)
	}

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	time.Sleep(3 * time.Second)
}
