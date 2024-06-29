package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	//go:embed data/prefixes.txt
	data []byte
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var prefixesMatchedList []string
	var resultPrefixesMatched []string
	var prefixesMaxLength int = 0
	fmt.Println("True caller start")

	workerMax := 2
	prefixes := bytes.Split(data, []byte{'\n'})
	prefixesLength := len(prefixes)
	fmt.Println("prefixesLength : ", prefixesLength)
	pages := prefixesLength / workerMax
	fmt.Printf("Pages: %v\n\n", pages)
	inputStr := "mDH" //"xyiuoXX" //xyi
	if workerMax == 1 {
		wg.Add(1)
		go fetchMatchedPrefix(&wg, 1, inputStr, &prefixes, &prefixesMatchedList, &prefixesMaxLength)
	} else if pages >= workerMax {
		// for counter :=0; counter< workerMax; counter++{

		// }
		wg.Add(workerMax)
		data_list1 := prefixes[0:pages]
		data_list2 := prefixes[pages:(prefixesLength - 1)]
		go fetchMatchedPrefix(&wg, 1, inputStr, &data_list1, &prefixesMatchedList, &prefixesMaxLength)
		go fetchMatchedPrefix(&wg, 2, inputStr, &data_list2, &prefixesMatchedList, &prefixesMaxLength)
	}

	// Code to measure
	duration := time.Since(start)

	wg.Wait()
	// fmt.Println("All Prefixes Found: ", prefixesMatchedList)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println("Max Length: ", prefixesMaxLength)
	fmt.Println(duration)

	findLongestMatches(&resultPrefixesMatched, &prefixesMatchedList , &prefixesMaxLength)
	fmt.Println("Final Prefixes List: ", resultPrefixesMatched)
}

// 350698
func fetchMatchedPrefix(wg *sync.WaitGroup, counter int, searchStr string, prefixes *[][]byte, prefixesMatchedList *[]string, prefixesMaxLength *int) {
	fmt.Println("Counter:", counter) 
	for _, prefix := range *prefixes {
		prefixString := string(prefix)
		if strings.HasPrefix(prefixString, searchStr) && len(prefixString) > 0 {
			*prefixesMatchedList = append(*prefixesMatchedList, prefixString)
			prefixLength := len(prefixString)
			if prefixLength > *prefixesMaxLength {
				*prefixesMaxLength = prefixLength
			} 
		}
	}
	wg.Done()
}

func findLongestMatches(resultPrefixesMatched *[]string, prefixesMatchedList *[]string, prefixesMaxLength *int){
	for _, prefix := range *prefixesMatchedList {
		if len(prefix) == *prefixesMaxLength{
			*resultPrefixesMatched = append(*resultPrefixesMatched, prefix)
		}
	}
}