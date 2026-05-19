package main

import "fmt"

func Generate(numChan chan int) {
	for i := 1; i <= 5; i++ {
		numChan <- i
	}
	// Important: close channel after sending all values
	close(numChan)
}

func main() {
	var numChan = make(chan int)

	go Generate(numChan)

	// range keeps receiving until channel is closed
	for v := range numChan {
		fmt.Println("Chan value:", v)
	}
	fmt.Println("All values received")
}
