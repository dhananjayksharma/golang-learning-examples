package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

var (
    //go:embed data/prefixes.txt
    data []byte
)

func main() { 
	start := time.Now() 
	fmt.Println("True caller start")
    // fmt.Println(string(data)) 
    fmt.Println("----------------------")
	workerMax := 5
	
    prefixes := bytes.Split(data, []byte{'\n'})
	prefixesLength := len(prefixes)
	fmt.Println("prefixesLength : ", prefixesLength)
	pages := prefixesLength / workerMax
	fmt.Printf("Pages: %v\n\n", pages)
	 
	inputStr := "XY" 
	fetchMatchedPrefix(1, inputStr, &prefixes)
  
	// Code to measure
	duration := time.Since(start)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)
	 
}

func fetchMatchedPrefix(counter int, searchStr string, prefixes *[][]byte){ 
	fmt.Println("Counter:", counter)
	for _, prefix := range *prefixes { 
		fmt.Println("all: ",string(prefix))
		if  strings.HasPrefix(searchStr, string(prefix)) && len(string(prefix))> 0{ 
			fmt.Println("Found: ",string(prefix))
		} 
    }
}