package main

import (
	"fmt"
)

func main() {
	arrData := []int{9, 1, 2, 3, 6}
	// arrData = []int{9, 4, 1, 2, 3, 6}
	// arrData = []int{1, 1, 3, 5, 2}
	// arrData = []int{-1, 4, 5}
	Solution(arrData)
}

func Solution(data []int) {
	firstSearch := true
	firstVal := findMinVal(data, 0, firstSearch)
	fmt.Println("firstVal:", firstVal)
	firstSearch = false
	secondVal := findMinVal(data, firstVal, firstSearch)
	fmt.Println("secondVal:", secondVal)
}
func findMinVal(data []int, search int, searchSeq bool) int {
	if searchSeq == false && search == 0 {
		return 0
	}
	out := 0
	for _, v := range data {
		if searchSeq {
			if out < v {
				out = v
			}
		} else {
			if search > v && out < v {
				out = v
			}
		}
		fmt.Println("out,v:", out, v)

	}
	fmt.Println("out:", out)
	return out
}
