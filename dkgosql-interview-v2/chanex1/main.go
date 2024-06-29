package main

import (
	"fmt"
	"sync"
)

func setChan(wg *sync.WaitGroup, defChan chan string, i string) {
	defer wg.Done()
	defChan <- i
}

var loopCounter = 0
var mu sync.Mutex

func main() {
	// var wg sync.WaitGroup		// option 1
	// var wg = sync.WaitGroup{}	// option 2
	wg := sync.WaitGroup{} // option 3
	data := []string{"a", "b", "c", "d", "e", "f", "g"}
	// defChan := make(chan string)
	defChan := make(chan string, 2)
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go setChan(&wg, defChan, data[i])
	}

	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go out(&wg, defChan)
	}

	wg.Wait()
}

func out(wg *sync.WaitGroup, defChan chan string) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	if loopCounter < 3 {
		loopCounter++
	}
	fmt.Printf("%s ", <-defChan)
	if loopCounter == 3 {
		loopCounter = 0
		fmt.Println("")
	}

}
