package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func double(in <-chan int) <-chan int {
	dbl := make(chan int)
	go func() {
		defer close(dbl)
		for n := range in {
			dbl <- n * 2
		}
	}()
	return dbl
}

func main() {
	for v := range double(generate(1, 2, 3, 6)) {
		fmt.Println("number:", v)
	}
}
