package main

import (
	"fmt"
	"time"
)

func main() {
	days := []int{1, 2, 3, 4, 5}
	for _, v := range days {
		go func(i int) {
			fmt.Println(" v:", i)
		}(v)
	}
	time.Sleep(3 * time.Second)
}
