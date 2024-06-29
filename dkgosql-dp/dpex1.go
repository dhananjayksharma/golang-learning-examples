package main

import "fmt"

func adder() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}
func fibo() func(int) int {
	result := make(map[int]int)
	return func(n int) int {
		// if _, ok := result[n]; ok {
		// 	return result[n]
		// } else {
		// 	if n < 2 {
		// 		return n
		// 	} else {
		// 		result[n] = fib(result[n-1]) + fib(result[n-2])
		// 		return result[n]
		// 	}
		// }
		result[n] = n
		fmt.Println(result)
		return result[n]
	}
}
func main() {
	fmt.Println("start")
	out := fibo()
	fmt.Println("out:", out(3))
}
