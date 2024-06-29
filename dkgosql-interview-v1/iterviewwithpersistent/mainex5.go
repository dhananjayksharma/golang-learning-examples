// Go program to illustrate the
// concept of select statement
package main

import "fmt"

// main function
func main() {

	// creating channel
	var mychannel chan bool

	select {
	case <-mychannel:

	default:
		fmt.Println("Not found")
	}

	mychannel <- true
}
