package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1) 
		println("Third one")
	}()
	go func() {
		msg1 := "hi"
		msg2 := "bye"
		defer wg.Done()
		fmt.Println("Message sent: " + msg1 + " and " + msg2)
		ch1 <- msg1
		ch1 <- msg2
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		rmsg1 := <-ch1
		rmsg2 := <-ch1
		println("Message received: " + rmsg1 + " and " + rmsg2)
	}()

	wg.Wait()
}