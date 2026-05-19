package main

/*
Problem
Create an unbuffered string channel. Launch a goroutine that sends "hello, channels" into it. In main, receive and print the value.
*/

import (
	"fmt"
)

func printdata(ch <-chan string) {
	ch <- "hello 11, channels"
}

func main() {

	ch2 := make(chan string)

	go printdata(ch2)

	msg2 := <-ch2

	fmt.Println("out 1:", msg2)

	// time.Sleep(2 * time.Second)
	fmt.Println("End of main")
}
