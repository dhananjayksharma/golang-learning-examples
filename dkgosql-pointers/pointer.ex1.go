package main

import "fmt"

func main() {
	var atikshAmount int = 100
	fmt.Printf("counter:%d, %p\n", atikshAmount, &atikshAmount)

	priyaAmount := &atikshAmount
	fmt.Printf("priyaAmount:%p, %p, %d\n", priyaAmount, &priyaAmount, *priyaAmount)

	*priyaAmount = 200
	fmt.Printf("priyaAmount val:%d\n", *priyaAmount)
	fmt.Printf("atikshAmount:%d, %p\n", atikshAmount, &atikshAmount)
	calculate(&atikshAmount)

	atikshAmount = 10
	fmt.Printf("after changes by saanvi priyaAmount val:%d\n", *priyaAmount)
	fmt.Printf("after changes by saanvi atikshAmount:%d, %p\n", atikshAmount, &atikshAmount)
}
func calculate(amount *int) {
	saanviAmount := amount
	fmt.Printf("saanviAmount val:%p, %d\n", saanviAmount, *saanviAmount)
	*saanviAmount = 10
	fmt.Printf("saanviAmount val:%p, %d\n", saanviAmount, *saanviAmount)
}
