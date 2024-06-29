package main

import "fmt"

type S struct {
	a, b, c string
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	fmt.Println(s, len(s), cap(s))
	newS := append(s, 55, 66)

	fmt.Println(newS, len(newS), cap(newS))
}
