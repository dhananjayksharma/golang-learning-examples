package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM!")
		case <-time.After(600 * time.Millisecond):
			fmt.Println("TIME TO QUIT!")
			return
			// default:
			// 	fmt.Println("   ...")
			// 	time.Sleep(100 * time.Millisecond)
		}
	}
}
