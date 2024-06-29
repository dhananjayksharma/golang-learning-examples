package main

import (
	"fmt"
)

func main() {

	// // Creating an array with inferred array length
	// tools := [...]string{"apm", "rum", "synthetic", "infra", "logs"}
	// fmt.Println("Original tools array(Before):", tools, ", len:", len(tools))

	// // Creating a new variable and initialize with tools
	// monitoring_tools := tools

	// fmt.Println("New monitoring tools array(before):", monitoring_tools)

	// // Change the value at index 0 to application performance monitoring
	// monitoring_tools[0] = "application performance monitoring"

	// fmt.Println("New monitoring tools array(After):", monitoring_tools)

	// fmt.Println("Original tools array(After):", tools)
	n := []int{1, 2, 3, 4, 5}
	out := Sum(n...)
	fmt.Println("sum of :", out)
}
func Sum(n ...int) int {
	sum := 0
	for v := range n {
		sum += v
	}
	return sum
}
