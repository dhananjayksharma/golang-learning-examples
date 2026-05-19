package main

import "fmt"

// Send 2 values on a channel, close it, then receive in a loop using the comma-ok idiom. Stop when the channel is closed and print whether each receive was valid.

func main() {
	var numChan = make(chan int, 2)

	numChan <- 10
	numChan <- 20

	close(numChan)

	for {
		v, ok := <-numChan
		if !ok {
			fmt.Printf("received: %d(ok=%v)\n", v, ok)
			fmt.Println("channel closed")
			break
		}
		fmt.Printf("received: %d(ok=%v)\n", v, ok)
	}

}
