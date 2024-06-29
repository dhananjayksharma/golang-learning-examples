package main

import "fmt"

func main() {
	var age int8
	age = 127

	fmt.Printf("Age:%v\n", age)

	age += 8
	fmt.Printf("Age:%v\n", age)

	age -= 127
	fmt.Printf("Age:%v\n", age)
}
