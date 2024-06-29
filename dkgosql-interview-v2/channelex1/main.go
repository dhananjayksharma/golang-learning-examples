package main

import (
	"fmt"
	"sync"
)

func processData(wg *sync.WaitGroup, input chan int) {
	defer wg.Done()
	m.Lock()
	defer m.Unlock()
	if loopCount < 3 {
		loopCount++
	}
	fmt.Printf("%d ", <-input)
	if loopCount == 3 {
		loopCount = 0
		fmt.Println("")
	}

}

var loopCount int
var wg sync.WaitGroup
var m sync.Mutex

func main() {
	loopCount = 0
	inputData := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chanData := make(chan int, 3)
	for i := 0; i < len(inputData); i++ {
		wg.Add(1)
		chanData <- inputData[i]
		go processData(&wg, chanData)
	}
	close(chanData)
	wg.Wait()
}
