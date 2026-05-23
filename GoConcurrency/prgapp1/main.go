package main

import (
	"fmt"
	"sync"
	"time"
)

func gorun(wg *sync.WaitGroup, from string) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(from, ":", i)
	}
}

func main() {

	ch1 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	// time.Sleep(1 * time.Second)

	go gorun(&wg, " gr 1")
	go gorun(&wg, " gr 2")

	// time.Sleep(1 * time.Second)
	// fmt.Println("done")

	// time.Sleep(5 * time.Second)
	wg.Wait()

	go func(s string) {
		ch1 <- s
	}("hello")

	fmt.Println("ch output:", <-ch1)
	close(ch1)
	// close(ch1)

	var newchan chan string
	<-newchan // <- "hello"

}
