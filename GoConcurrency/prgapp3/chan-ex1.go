package main

import "fmt"

func qube(ch2 chan<- int, n int) {
	defer close(ch2)
	fmt.Printf("qube n: %v\n", n)
	v := n * n * n
	ch2 <- v
}

func square(ch1 chan<- int, n int) {
	fmt.Printf("Square n: %v\n", n)
	defer close(ch1)
	v := n * n
	ch1 <- v
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	for i := 0; i < 5; i++ {
		fmt.Println("Loop i:", i)
		go square(ch1, i)
		go qube(ch2, i)
	}

	select {
	case v1, ok := <-ch1:
		if ok == false {
			fmt.Println("ch1 Close ", ok)
			break
		}
		fmt.Printf("Square: %v\n", v1)
	case v2, ok := <-ch2:
		if ok == false {
			fmt.Println("ch1 Close ", ok)
			break
		}
		fmt.Printf("Qube: %v\n", v2)
	}

}
