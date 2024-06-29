package main

import (
	"fmt"
)

func getOccurance(data []int) {
	out := make(map[int]int)
	for _, v := range data {
		if _, ok := out[v]; ok {
			out[v] += 1
		} else {
			out[v] = 1
		}
	}
	fmt.Println("out:", out)
	var output []int
	for _, v := range data {
		val := out[v]
		if val > 1 {
			output = append(output, v)
		}
	}
	fmt.Println("output:", output)
}
func main() {
	fmt.Println("main")
	var data = []int{3, 2, 1, 2, 1, 4, 5, 8, 6, 7, 4, 2, 8}
	getOccurance(data)
}
