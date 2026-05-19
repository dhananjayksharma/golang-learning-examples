package main

import (
	"fmt"
)

func generator(n int, chsq chan<- int) {
	for i := 1; i <= n; i++ {
		chsq <- i * i
	}
	close(chsq)
}

func main() {
	chsq := make(chan int)

	go generator(3, chsq)

	for v := range chsq {
		fmt.Println("value :", v)
	}

}
