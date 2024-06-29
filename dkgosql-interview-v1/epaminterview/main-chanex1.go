package main

import "fmt"

func main() {
	fmt.Println("main start")
	a := 0

	waitchan := make(chan int)
	done := make(chan int)
	go func(waitchan chan int) {
		a++
		waitchan <- 1
	}(waitchan)

	go func(waitchan, done chan int) {
		<-waitchan
		done <- 1

		a *= 2
		waitchan <- 1
	}(waitchan, done)

	<-done
	<-waitchan
	fmt.Println("a:", a)

}
