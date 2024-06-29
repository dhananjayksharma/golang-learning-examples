package main

import (
	"fmt"
	"time"
)

func main() {

	// For loop
	for i := 1; i < 6; i++ {

		// Prints these util loop stops
		fmt.Println("****Welcome to GeeksforGeeks***")
		fmt.Println("A CS-Portal!")
	}

	// Select statement
	select {

	// Using case statement to receive
	// or send operation on channel and
	// calling After() method with its
	// parameter
	case <-time.After(3 * time.Second):

		// Printed when timed out
		fmt.Println("Time Out!")
	}
}
