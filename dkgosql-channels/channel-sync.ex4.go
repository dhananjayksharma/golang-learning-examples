package main

import "fmt"

func main() {
	datachan := make(chan int)
	waitchain := make(chan struct{})

	go printData(datachan, waitchain)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d", i)
		datachan <- i
		<-waitchain
	}
}

func printData(datachan chan int, waitchain chan struct{}) {
	for v := range datachan {
		fmt.Printf("%d\n", v)
		waitchain <- struct{}{}
	}
}
