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
		fmt.Println("In go routing 1a")
		ch <- "hello 1, channels" // blocks until receiver ready
		fmt.Println("In go routing 1b")
	}()

	go func() {
		fmt.Println("In go routing 2a")
		ch <- "hello 2, channels" // blocks until receiver ready
		fmt.Println("In go routing 2b")
	}()

	msg := <-ch // receives, unblocks goroutine
	fmt.Println("out 1:", msg)

	msg1 := <-ch // receives, unblocks goroutine
	fmt.Println("out 2:", msg1)
	time.Sleep(2 * time.Second)
	fmt.Println("End of main")
}
