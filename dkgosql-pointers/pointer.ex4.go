package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	a := &computer{"Apple"}
	b := a
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	change(b)
	fmt.Printf("%v\n", *a)
	change(b)
	fmt.Printf("%v\n", *b)
}

func change(c *computer) {
	c.brand = "Indie"
	// c = nil
	fmt.Printf("%p", c)
}
