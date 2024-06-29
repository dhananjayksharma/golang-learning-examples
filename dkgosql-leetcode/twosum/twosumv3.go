package main

import (
	"fmt"
)

func Sum(input []int, target int) (int, bool) {
	output := make(map[int]bool)
	for i := 0; i < len(input); i++ {
		currVal := input[i]
		remain := target - currVal
		if _, ok := output[remain]; ok {
			return currVal + remain, true
		} else {
			output[currVal] = true
		}
	}
	return 0, false
}
func main() {
	input := []int{2, 3, 4, 5, 6, 7, 8}
	target := 6
	out, found := Sum(input, target)
	if found {
		fmt.Println("Found:", out)
	} else {
		fmt.Println("Not Found:", out)
	}
}
