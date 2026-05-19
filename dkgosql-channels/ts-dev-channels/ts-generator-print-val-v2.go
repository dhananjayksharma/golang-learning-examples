package main

import (
	"fmt"
)

func square(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i * i
		}
	}()
	return ch
}

func main() {
	out := square(5)
	// for v := range out {
	// 	fmt.Println(v)
	// }
	for {
		v, ok := <-out
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
