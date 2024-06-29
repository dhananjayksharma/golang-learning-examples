package main

import (
	"fmt"
	"sync"
)

func g1(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- 42
}

func g2(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- 43
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go g1(&wg, ch1)
	go g2(&wg, ch2)
	// time.Sleep(2 * time.Second)
	select {
	case v1 := <-ch1:
		fmt.Println("Got ch1: ", v1)
	case v2 := <-ch2:
		fmt.Println("Got ch2: ", v2)
	default:
		fmt.Println("The default case!")
	}
	wg.Wait()
	// time.Sleep(2 * time.Second)
}
