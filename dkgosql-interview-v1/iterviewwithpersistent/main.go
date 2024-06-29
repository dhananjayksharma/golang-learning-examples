package main

import (
	"fmt"
	"strings"
	"sync"
)

func upper(wg *sync.WaitGroup, done chan bool, str string) {
	defer wg.Done()
	//str := "abcdefghijklm"
	for _, v := range str {
		// fmt.Println("Waiting for time to elapsed, and to push to lower chan")
		// time.Sleep(1 * time.Second)
		done <- true
		fmt.Printf("%v", strings.ToUpper(string(v)))
		<-done
	}

}

func lower(wg *sync.WaitGroup, done chan bool, str string) {
	defer wg.Done()
	for _, v := range str {
		<-done
		fmt.Printf("%v", strings.ToLower(string(v)))
		// fmt.Println("Waiting for time to elapsed and push done channel to upper")
		// time.Sleep(1 * time.Second)
		done <- true
	}
}

func main() {
	var done = make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	str := "abcdefghijklm"
	go upper(&wg, done, str)
	go lower(&wg, done, str)

	wg.Wait()
}
