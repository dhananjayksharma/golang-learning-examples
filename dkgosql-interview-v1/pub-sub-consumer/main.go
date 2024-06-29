package main

import (
	"fmt"
	"sync"
)

func producer(wg *sync.WaitGroup, table chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		table <- i * 5
	}
	close(table)
}

func consumer(wg *sync.WaitGroup, table chan int, consid int) {
	defer wg.Done()
	for out := range table {
		fmt.Printf("read by cons#%d : %d\n", consid, out)
	}
}

func main() {
	var wg sync.WaitGroup
	var table = make(chan int)
	wg.Add(1)
	go producer(&wg, table)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		consid := i + 1
		go consumer(&wg, table, consid)
	}
	wg.Wait()
	fmt.Println("main done")
}
