package main

import (
	"fmt"
	"time"
)

func dataProcess(inputData []int, ch chan int) {
	for i := 0; i < 100; i++ { //range inputData {
		ch <- i
		time.Sleep(10 * time.Millisecond)
	}
	close(ch)
}

func main() {
	//     create an unbuffered channel
	ch := make(chan int, 90)
	inputData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go dataProcess(inputData, ch)
	for d := range ch {
		fmt.Println(d)
	}
	fmt.Println("here")
}
