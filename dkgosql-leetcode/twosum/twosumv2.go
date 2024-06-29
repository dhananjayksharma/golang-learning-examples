package main

import (
	"fmt"
	"sort"
)

func Sum(input []int, target int) (int, bool) {
	sort.Ints(input)
	fmt.Println("input:", input)
	sum := 0
	left := 0
	right := len(input) - 1
	for left < right {
		sum = input[left] + input[right]
		if sum == target {
			fmt.Println("found these here:", left, right, "Val:", input[left], input[right])
			return sum, true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return 0, false
}
func main() {
	input := []int{2, 13, 19, 23, 4, 3, 53, 8}
	target := 13
	out, found := Sum(input, target)
	if found {
		fmt.Println("Found:", out)
	} else {
		fmt.Println("Not Found:", out)
	}
}
