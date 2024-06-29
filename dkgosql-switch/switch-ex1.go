package main

import (
	"fmt"
	"sync"
)

func fibo(c chan int, quit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case d := <-quit:
			fmt.Printf("Quit data%d\n", d)
			close(c)
			return
		}
	}
}
func main() {
	fmt.Println("Start main")
	c := make(chan int)
	quit := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("Out val: %d\n", <-c)
		}
		quit <- struct{}{}
	}(&wg)
	go fibo(c, quit, &wg)

	wg.Wait()
}
