package main

import "fmt"

func factorial(n int) int {
	fmt.Printf("n=%d\n", n)
	if n == 0 {
		return 1
	}

	out := n * factorial(n-1)
	return out
}
func main() {
	var n int = 5
	out := factorial(n)
	fmt.Printf("Factorial %d = %d", n, out)
}
