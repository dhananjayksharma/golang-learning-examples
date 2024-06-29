package main

import "fmt"

func adder() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}
func main() {
	a := adder()
	fmt.Println("a 1:", a(1))
	fmt.Println("a 2:", a(2))
	fmt.Println("a 3:", a(3))

	b := adder()
	fmt.Println("b 1:", b(10))
	fmt.Println("b 2:", b(20))
	fmt.Println("b 3:", b(30))
}
