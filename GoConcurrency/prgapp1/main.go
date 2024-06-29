package main

import (
	"fmt"
	"sync"
	"time"
)

func f(from string) { 
    for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
        fmt.Println(from, ":", i) 
    }
}

func main() {

	// ch1 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

    go f(" gr 1")
	
	go f(" gr 2")
 

    time.Sleep(time.Second)
    fmt.Println("done")
	wg.Wait()
}