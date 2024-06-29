package main

import (
	"fmt"
	"reflect"
)

type S struct {
	a, b, c string
}

func main() {
	// option 1
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	fmt.Println(x == y)
	// option 2
	var a interface{}
	a = &S{"a", "b", "c"}
	var b interface{}
	b = &S{"a", "b", "c"}
	fmt.Println(a == b, a, b)

	var p interface{}
	p = []int{1, 2}
	var m interface{}
	m = []int{1, 2}
	fmt.Println(reflect.DeepEqual(p, m))
}
