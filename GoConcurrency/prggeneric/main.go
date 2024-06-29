package main

import "fmt"

func sum[T int | float64](a T, b T) T {
	s := a + b

	return s
}

func main() {
	var out interface{}
	out = sum(3, 4)
	res, ok := out.(int)
	if ok {
		fmt.Println("Int:", res)
	}
	res2, ok := out.(float64)
	if ok {
		fmt.Println("float64:", res2)
	}

}
