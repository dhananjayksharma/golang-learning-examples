package main

import (
	"fmt"
	"sync"
)

var data int

func main() {
	var mu sync.Mutex
	var chanDone = make(chan struct{}, 2)
	go func(chanDone chan struct{}, mu *sync.Mutex) {
		mu.Lock()
		defer mu.Unlock()
		data++
		fmt.Printf("Step 1 the value is %v.\n", data)
		chanDone <- struct{}{}
	}(chanDone, &mu)
	<-chanDone

	if data == 0 {
		fmt.Printf("Step 2 the value is %v.\n", data)
	}

	fmt.Printf("Step 3 the value is %v.\n", data)

}

// 1
// 2
// 3
