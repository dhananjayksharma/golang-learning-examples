package main

import "fmt"

func SendData(datachan chan int) {
	for i := 1; i <= 10; i++ {
		datachan <- i
	}
	close(datachan)
}

func main() {
	var datachan = make(chan int)
	go SendData(datachan)
	i := 21
	for d := range datachan {
		fmt.Printf("m%d", i)
		i++
		fmt.Printf("c%d ", d)
	}
}
