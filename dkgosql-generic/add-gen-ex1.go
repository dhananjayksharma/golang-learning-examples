package main

import "fmt"

func add[K int | float64](a K, b K) K {
	return a + b
}
func main() {
	out := add(1.3, 3.4)
	// out := add(1, 3)
	fmt.Println("Sum:", out)
}
