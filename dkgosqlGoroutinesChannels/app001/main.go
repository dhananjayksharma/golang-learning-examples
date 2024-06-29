package main

import (
	"fmt"
)
 
func main(){ 
	fmt.Println("Main started...")
	//create channel of int
	r := make(chan int, 1)

	w := make(chan int, 1)

	// call to goroutine
	go readChannelData(r)
	for {
        select {
        case msg1 := <-r:
			writeChannelData(w, msg1) 
		}
    } 
   	// fmt.Println("Main ended...")
}
 
func readChannelData(c chan<- int) {
	c<- 10
	fmt.Println("Setting data readChannelData")
 }

 func writeChannelData(w <-chan int, val int) { 
	fmt.Println("Data in writeChannelData: ", val)
	
 }