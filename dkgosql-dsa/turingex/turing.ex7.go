package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float64
}
type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.height * r.width
}

func main() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Printf("%+v, %v\n", s, reflect.TypeOf(s))
	fmt.Printf("%+v, %v\n", r, reflect.TypeOf(s))
	fmt.Println(s == r)
}
