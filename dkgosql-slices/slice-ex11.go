package main

import "fmt"

func add(a []int) {
	a[1] = 90
	// a = append(a, 99)
	// fmt.Println("In add:", a)
}

func main() {
	// a := make([]int, 2) // [0, 0]
	// add(a)

	// fmt.Println("Outer:", a) // [0, 90]

	b := make([]int, 1, 2) //[0]
	fmt.Println("Len, Cap", len(b), cap(b))
	b = append(b, 90)
	fmt.Println("Len, Cap", len(b), cap(b))
	b = append(b, 92)
	fmt.Println("Len, Cap", len(b), cap(b))
	// add(b)

	fmt.Println("Outer:", b)
}
