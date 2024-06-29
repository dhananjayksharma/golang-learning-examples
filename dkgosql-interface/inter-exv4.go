package main

import "fmt"

func main() {
	fmt.Println("start")
	var dataStream chan interface{}
	close(dataStream)
	// dataStream <- struct{}{}
	// fmt.Println("dataStream:", dataStream)
	// fmt.Println("read dataStream:", <-dataStream)
	// dataStream = make(chan interface{}, 4)
	// dataStream <- 23
	// fmt.Println("dataStream:", dataStream, cap(dataStream), len(dataStream), <-dataStream)
}
