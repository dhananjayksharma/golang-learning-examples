package main

/*
Problem
Create an unbuffered string channel. Launch a goroutine that sends "hello, channels" into it. In main, receive and print the value.
*/

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // unbuffered

	go func() {
		fmt.Println("Sender line 1a")
		ch <- "hello 1, channels" // blocks until receiver ready
		fmt.Println("Sender line 1b")
	}()

	go func() {
		fmt.Println("receiver line 1a")
		fmt.Println("Here:", <-ch)
		fmt.Println("receiver line 1b")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("End of main")
}
