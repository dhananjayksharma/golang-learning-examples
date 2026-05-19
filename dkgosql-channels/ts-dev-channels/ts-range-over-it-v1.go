package main

/*
Problem
Send 5 integers from a goroutine into a channel, then close it. In main, use range to receive all values without knowing the count.
*/

import (
	"fmt"
)

func producer(num chan<- int) {
	for i := 0; i < 5; i++ {
		num <- (i + 1*200)
	}

	close(num) // signals no more values

}

func main() {
	num := make(chan int, 5) // unbuffered
	go producer(num)

	for v := range num { // exits when num is closed & drained
		fmt.Printf("value: %v\n", v)
	}

	fmt.Println("End of main")
}

/*
range on a channel receives until it's closed. Only the sender should close a channel. Closing twice panics. Never close from the receiver side.
*/
