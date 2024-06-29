package main

import (
	"fmt"
	"runtime"
)

func squares(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	c := make(chan int, 3)
	go squares(c)

	fmt.Print("NG1:", runtime.NumGoroutine(), "\n")

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Print("NG2:", runtime.NumGoroutine(), "\n")

}
