package main

import "fmt"

const N = 4

func main() {
	data := make(map[int]*int)
	for i := 1; i <= N; i++ {
		data[i] = &i
	}
	// fmt.Printf("%v", data)

	for k, v := range data {
		fmt.Printf("%d=%v\n", k, *v)
	}

	/*
	   1:5
	   2:5
	   3:5
	   4:5
	*/
	// fmt.Printf("%p\n", data)

	// b := 90
	// fmt.Printf("%p = %d\n", &b, b)

	// b = 23
	// fmt.Printf("%p = %d\n", &b, b)
}
