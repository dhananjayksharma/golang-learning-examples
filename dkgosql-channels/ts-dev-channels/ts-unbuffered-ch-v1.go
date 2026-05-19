package main

/*
Problem
Create an unbuffered string channel. Launch a goroutine that sends "hello, channels" into it. In main, receive and print the value.
*/

import (
	"fmt"
)

func main() {
	ch := make(chan string) // unbuffered
	go func() {
		ch <- "hello 1, channels" // blocks until receiver ready
	}()

	msg := <-ch // receives, unblocks goroutine
	fmt.Println("out 1:", msg)

	// time.Sleep(2 * time.Second)
	fmt.Println("End of main")
}
