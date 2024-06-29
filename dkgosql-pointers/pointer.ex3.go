package main

import (
	"fmt"
)

func main() {
	type computer struct {
		brand string
	}

	var a, b *computer
	fmt.Println(a == b, a == nil)

	fmt.Printf("a=%p and b=%p\n", a, b)

	a = &computer{"Apple"}
	b = &computer{"Apple"}

	fmt.Printf("a=%p and b=%p\n", &a, &b)
	fmt.Printf("a=%v and b=%v\n", *a, *b)
	fmt.Print(" ", a == b, " ", *a == *b)
}
