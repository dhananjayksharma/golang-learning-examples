package main

import (
	"fmt"
)

//Problem
// Write two functions: send accepts a send-only channel and sends a string.
// receive accepts a receive-only channel and prints what it gets. Use directional types.

func producer(number chan<- string) { // send-only
	for i := 0; i < 5; i++ {
		number <- fmt.Sprintf("platform engineering %d", i+1*3)
	}
	close(number)
}

func consumer(number <-chan string, waitCh chan struct{}) { // recieve-only
	for v := range number {
		fmt.Println("value:", v)
	}
	waitCh <- struct{}{}
	close(waitCh)
}

func main() {
	var number = make(chan string, 2)
	var waitCh = make(chan struct{})
	go consumer(number, waitCh)

	go producer(number) // bidirectional auto-converts

	<-waitCh
	fmt.Println("end of main")
}

// Key Concept
// Directional channels enforce intent at compile time. chan<- is write-only; <-chan is read-only. Bidirectional channels implicitly convert to directional — but not back.
