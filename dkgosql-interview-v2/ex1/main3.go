package main

import (
	"fmt"
	"sync"
	"time"
)

var data int

func main() {
	var mu sync.Mutex
	var chanDone = make(chan struct{})
	go func(chanDone chan struct{}, mu *sync.Mutex) {
		mu.Lock()
		defer mu.Unlock()
		data++
		fmt.Printf("Step 1 the value is %v.\n", data)
		chanDone <- struct{}{}
	}(chanDone, &mu)
	<-chanDone
	time.Sleep(1 * time.Second)
	mu.Lock()
	if data == 0 {
		data += 3
		fmt.Printf("Step 2 the value is %v.\n", data)
	}
	mu.Unlock()

	fmt.Printf("Step 3 the value is %v.\n", data)

}

// 1
// 2
// 3
