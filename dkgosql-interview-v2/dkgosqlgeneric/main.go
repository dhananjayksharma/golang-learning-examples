package main

import (
	"fmt"
)

func sum[K int | float64 | int64 | string](a K, b K, c K) K {
	aaa := a + b + c
	return aaa
}
func main() {
	out := sum(1, 4, 5)
	fmt.Println("out:", out)
}
