package main

import (
	"fmt"
)

func tryReceive(msgChan <-chan string) {
	select {
	case v := <-msgChan:
		fmt.Println("got msg:", v)
	default:
		fmt.Println("No message")
	}
}

func main() {
	var msgChan = make(chan string, 1)

	tryReceive(msgChan)
	msgChan <- "Hello"
	tryReceive(msgChan)

	// time.Sleep(2 * time.Second)
}
