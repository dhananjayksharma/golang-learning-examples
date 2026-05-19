package main

import (
	"errors"
	"fmt"
)

func worker(id int, done chan<- struct{}, errChan chan error) {
	fmt.Printf("worker %d done\n", id)
	if id > 1000 {
		errChan <- errors.New("id must be less than 1000")
	}
	errChan <- errors.New("")
	done <- struct{}{}
}

func main() {
	var errChan = make(chan error)
	var done = make(chan struct{}, 3)
	for i := 1; i <= 3; i++ {
		go worker(i, done, errChan)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("error", <-errChan)
		<-done
	}

	fmt.Println("all done")
}
