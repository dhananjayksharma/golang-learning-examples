package main

import "fmt"

func Sum(input []int, target int) (int, bool) {
	out := 0
	if len(input) < 2 {
		return out, false
	}
	if target == 0 {
		return out, false
	}
	for i, v1 := range input {
		for i2, v2 := range input {
			if v1+v2 == target && i != i2 {
				fmt.Println("found these:", v1, v2)
				return target, true
			}
		}
	}
	return out, false
}
func main() {
	input := []int{-2, -7, 9, 23, 4, 3, 5, 8}
	target := -9
	out, found := Sum(input, target)
	if found {
		fmt.Println("Found:", out)
	} else {
		fmt.Println("Not Found:", out)
	}
}
