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
	
	prefixes := bytes.Split(data, []byte{'\n'})
	prefixesLength := len(prefixes)
	fmt.Println("prefixesLength : ", prefixesLength)
	
	workerMax :=10
	// Find Remainder
	remainder := prefixesLength % workerMax 
	fmt.Println("remainder: ", remainder)
	pages := prefixesLength / workerMax
	// fmt.Printf("Pages: %v\n\n", pages)
	inputStr := "xyi" //"xyiuoXX" //xyi
	if workerMax == 1 {
		wg.Add(1)
		go fetchMatchedPrefix(&wg, 1, inputStr, &prefixes, &prefixesMatchedList, &prefixesMaxLength)
	} else if pages >= workerMax {
		skip, take := 0, pages
		for counter := 0; counter < workerMax; counter++ {
			wg.Add(1)
			data_list := prefixes[skip:take]
			fmt.Println("Page from :", skip, take)
			skip = skip + pages
			take = take + pages
			go fetchMatchedPrefix(&wg, counter, inputStr, &data_list, &prefixesMatchedList, &prefixesMaxLength)
		}

		if remainder > 0 {
			wg.Add(1)
			skip = pages*workerMax
			take = pages*workerMax+remainder
			data_list := prefixes[skip : take]
			fmt.Println("Page from :", skip, take)
			go fetchMatchedPrefix(&wg, workerMax, inputStr, &data_list, &prefixesMatchedList, &prefixesMaxLength)
		}
	}

	// Code to measure
	duration := time.Since(start)

	wg.Wait()
	fmt.Println("Max Length: ", prefixesMaxLength)
	fmt.Println(duration)

	findLongestMatches(&resultPrefixesMatched, &prefixesMatchedList, &prefixesMaxLength)
	fmt.Println("Final Prefixes List: ", resultPrefixesMatched)
}

func fetchMatchedPrefix(wg *sync.WaitGroup, counter int, searchStr string, prefixes *[][]byte, prefixesMatchedList *[]string, prefixesMaxLength *int) {
	fmt.Println("Counter:", counter)
	for _, prefix := range *prefixes {
		prefixString := string(prefix)
		// fmt.Println("prefix:", string(prefix))
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

func findLongestMatches(resultPrefixesMatched *[]string, prefixesMatchedList *[]string, prefixesMaxLength *int) {
	for _, prefix := range *prefixesMatchedList {
		if len(prefix) == *prefixesMaxLength {
			*resultPrefixesMatched = append(*resultPrefixesMatched, prefix)
		}
	}
}
