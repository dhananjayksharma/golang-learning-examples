package main

/*
Problem
Create a buffered int channel with capacity 3.
Send 3 integers without a goroutine (no blocking).
Then receive and print all 3.
*/

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3) // buffered: holds 3 items

	ch <- 100 // non-blocking (buffer has space)
	ch <- 101
	ch <- 102
	close(ch)
	// ch <- 40 //→ deadlock! buffer full, no receiver

	fmt.Printf("chan :%v\n", <-ch)
	fmt.Printf("chan :%v\n", <-ch)
	fmt.Printf("chan :%v\n", <-ch)
	// fmt.Printf("chan :%v\n", <-ch)

	fmt.Println("End of main")

	/*
		Key Concept
		Buffered channels decouple sender and receiver up to capacity. Useful when producer is faster than consumer in bursts. Sending to a full buffer blocks; receiving from empty blocks.
	*/
}
