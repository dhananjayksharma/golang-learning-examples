package main

import (
	"fmt"
	"time"
)

func print(i int) {
	fmt.Println(i)
}

func main() {

	for i := 0; i < 10; i++ {

		go print(i)
	}

	time.Sleep(10 * time.Second)
}
