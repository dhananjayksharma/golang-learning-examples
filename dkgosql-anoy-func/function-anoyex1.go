package main

import "fmt"

func main() {
	sum := func(a, b int) int { // anoynomus
		return a + b
	}

	func(a, b int) { // Immediate
		fmt.Println("Immediate:", a+b)
	}(1, 2)

	fmt.Println("later:", sum(3, 4))
}
