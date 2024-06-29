package main

import "fmt"

// fibonacchi sequence
//
//	0, 1, 1, 2, 3, 5, 8, 13, 21

func getFibonacci(num int) []int {
	output := make([]int, 0, 10)
	output = append(output, 0, 1)
	fmt.Println(output)
	for i := 1; i < num; i++ {
		sum := output[i-1] + output[i]
		output = append(output, sum)
	}
	return output
}
func main() {
	num := 8
	out := getFibonacci(num)

	fmt.Println("Output:", out)
}
