package main

import "fmt"

func main() {
	var data = make(chan bool)
	go func() {
		for {
			select {
			case v := <-data:
				fmt.Println("Break:", v)
				return
			default:
				fmt.Println("its default")
			}
		}
	}()
	data <- true
	// fmt.Println("line 16")
	// data <- true
	// fmt.Println("line 18")
}
