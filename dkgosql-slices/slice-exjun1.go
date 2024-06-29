package main

import (
	"fmt"
)

func main() {
	fmt.Println("start main")
	var rainhistory []int
	if rainhistory == nil {
		fmt.Printf("rainhistory is %T\n", rainhistory)
	}
	fmt.Printf("rainhistory is %d, %d\n", len(rainhistory), cap(rainhistory))
	rainhistory = append(rainhistory, 11)

	fmt.Printf("rainhistory is %v\n", rainhistory)
	fmt.Printf("rainhistory is %d, %d\n", len(rainhistory), cap(rainhistory))

	// Empty slice
	var slice = []string{}

	if slice == nil {
		fmt.Printf("slice is %T\n", slice)
	}
	// or
	var slice1 = make([]string, 0)
	if slice1 == nil {
		fmt.Printf("slice1 is %T\n", slice1)
	} else {
		fmt.Printf("slice1 is not null %T\n", slice1)
	}
}
