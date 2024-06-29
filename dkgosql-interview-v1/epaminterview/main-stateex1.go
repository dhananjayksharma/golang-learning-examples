package main

import "fmt"

func main() {
	a := make(chan int, 1)
	waitchan := make(chan int)
	go func(a, waitchan chan int) {
		waitchan <- 1
		a <- 1444
	}(a, waitchan)

	go func(a, waitchan chan int) {

		for out := range a {
			fmt.Println(out)
		}
		<-waitchan
		waitchan <- 1
	}(a, waitchan)

	<-waitchan
}
