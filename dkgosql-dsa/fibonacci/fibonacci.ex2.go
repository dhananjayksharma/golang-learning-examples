package main

import "fmt"

// fibonacchi sequence
//
//	1 1 2 3 5 8 13 21 34 55

func getFibonacci2() [10]int {
	output := [10]int{}
	output[0] = 1
	output[1] = 1
	fmt.Println(output)
	for i := 2; i < 10; i++ {
		output[i] = output[i-1] + output[i-2]
	}
	return output
}
func main() {
	out := getFibonacci2()

	fmt.Println("Output:", out)
}
