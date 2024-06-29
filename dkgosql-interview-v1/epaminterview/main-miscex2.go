package main

import "fmt"

type X struct {
}

type Y struct {
}

func (y *Y) Do() int {
	return 6
}

func main() {
	var y = Y{}
	fmt.Println(y.Do())
}
