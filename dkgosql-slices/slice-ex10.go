package main

import "fmt"

func add(a []int) {
	a[1] = 90
	a = append(a, 99)
	fmt.Println("In add:", a)
}

func main() {
	a := []int{1, 2, 3}
	add(a)

	fmt.Println("Outer:", a)
}
