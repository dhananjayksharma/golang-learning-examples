package main

import "fmt"

type X struct {
}

func (x X) Y() Y {
	return Y{}
}
func (x X) W() int {
	return 4
}

type Y struct {
	X
}

func (y *Y) W() int {
	return 6
}

func main() {
	var y Y
	fmt.Println(y.W())
}
