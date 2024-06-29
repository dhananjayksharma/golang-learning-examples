package main

import "fmt"

type E interface {
	~int
}
type Point []int32

func Scale[S ~[]E, E int](s S, sc E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * sc
	}
	return r
}
func main() {
	fmt.Println("start")
	var e  E = {1, 2, 5, 6}
	out := Scale(e, 23)
	fmt.Println(out)
}
