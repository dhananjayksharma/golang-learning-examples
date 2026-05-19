package main

import (
	"fmt"
)

func sender(msg chan<- string) { // send only
	msg <- "Platform engineering"
	close(msg) // close channel
}

func reader(msg <-chan string, waitCh chan struct{}) { // read only
	fmt.Printf("Reading: %s\n\n", <-msg)
	waitCh <- struct{}{}
}

func main() {
	var msg = make(chan string, 1)
	var waitCh = make(chan struct{})
	go sender(msg)

	go reader(msg, waitCh)

	<-waitCh
	fmt.Println("End of main")
}
