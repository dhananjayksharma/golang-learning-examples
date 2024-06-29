package main

import "fmt"

func getHash(data []int, target int) (int, bool) {
	dataHash := make(map[int]bool)
	for _, v := range data {
		if _, ok := dataHash[v]; ok { // && v == target
			return v, true
		} else {
			dataHash[v] = true
		}

	}
	return 0, false
}
func main() {
	data := []int{2, 6, 5, 1, 3, 4}
	target := 2
	num, out := getHash(data, target)
	if out {
		fmt.Println("found:", out, num)
	} else {
		fmt.Println("not found:", out, num)
	}
}
