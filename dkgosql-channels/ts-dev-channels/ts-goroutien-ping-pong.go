package main

import (
	"fmt"
	"time"
)

func A(name string, in <-chan string, out chan<- string) {
	for i := 0; i <= 3; i++ {
		msg := <-in
		fmt.Printf("%s recived: %s\n", name, msg)
		out <- name
	}
}

func main() {
	in := make(chan string)
	out := make(chan string)

	go A("pong", in, out)
	in <- "start"
	A("ping", out, in)
	time.Sleep(2 * time.Second)

}
