package main

import (
	"fmt"
	"time"
)

func worker(done chan<- struct{}) {
	for i := 0; i < 10; i++ {
		waitTime := time.Duration(i)
		fmt.Println("processing time is ", waitTime)
		time.Sleep(waitTime * time.Second)
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{}, 1)

	go worker(done)
	fmt.Println("waiting for status")
	<-done
	fmt.Println("done")
}
