package main

import (
	"fmt"
)

func main() {
	// Creating a channel
	// Using make keyword
	channel := make(chan int)
	var quit = make(chan bool)

	go func(channel chan int, quit chan bool) {
		for i := 0; i < 10000; i++ {
			channel <- i
		}
		close(channel)
		quit <- true
	}(channel, quit)

	for v := range channel {
		fmt.Println(v)
	}
	<-quit
}
