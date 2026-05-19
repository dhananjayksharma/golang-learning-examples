package main

import (
	"fmt"
	"time"
)

func slowOp() <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second) //Slower than timeout
		ch <- "result"
	}()

	return ch
}

func payment(){
	
}

func main() {

}
