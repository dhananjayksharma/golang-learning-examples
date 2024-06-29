package main

import (
	"fmt"
	"time"
)

func SendData(datachan chan int) {
	for i := 1; i <= 10; i++ {
		datachan <- i
	}
	close(datachan)
}

func main() {
	startTime := time.Now()
	var datachan = make(chan int, 4) // blocking
	go SendData(datachan)

	for i := 1; i <= 10; i++ {
		fmt.Printf("m%d", i)
		fmt.Printf("c%d\n", <-datachan)
	}
	endTime := time.Now().Sub(startTime)
	fmt.Printf("Total time:%d", endTime.Microseconds())
}
