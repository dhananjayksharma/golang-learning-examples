package main

import (
	"fmt"
)

func recovery(done chan bool) {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r) // step 2
	}
	done <- true
}

func sum(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b) // step 1
	done := make(chan bool)
	go divide(a, b, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	defer recovery(done)
	fmt.Printf("%d / %d = %d", a, b, a/b)

}

func main() {
	sum(5, 3)
	fmt.Println("normally returned from main") // step 3
}
