package main

import (
	"fmt"
)

func main() {
	apply := func(a, b int, op func(int, int) int) int { return op(a, b) }

	multiply := func(a, b int) int { return a * b }

	fmt.Println(apply(3, 4, multiply))           // 12
	fmt.Println(apply(3, 4, func(a, b int) int { // inline anonymous
		return a + b
	}))

	add := func(a, b int) int {
		fmt.Println("In exec")
		return a + b
	}

	fmt.Println(apply(6, 5, add))
}
