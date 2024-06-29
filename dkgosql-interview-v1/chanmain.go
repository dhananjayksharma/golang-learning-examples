package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	dataChan := make(chan int, 3)
	data := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(data); i++ {
		wg.Add(2)
		go SetData(&wg, dataChan, data[i])
		go GetData(&wg, dataChan)
	}

	wg.Wait()
}

func SetData(wg *sync.WaitGroup, dataChan chan int, i int) {
	defer wg.Done()
	dataChan <- i * 2
}

func GetData(wg *sync.WaitGroup, dataChan chan int) {
	defer wg.Done()
	fmt.Println("Value read:", <-dataChan)
}

// diff
//array vs slice	-->
