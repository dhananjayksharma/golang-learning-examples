package main

import "fmt"

func main() {
	out := getFib(5)
	for i, v := range out {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", i, v)
	}
}

func getFib(n int) []int {
	a := 1
	b := 1
	fib := 0
	var out = make([]int, 0, n)
	for i := 0; i < n; i++ {
		fib = fib + a
		a, b = b, fib
		out = append(out, a)
	}
	return out
}
