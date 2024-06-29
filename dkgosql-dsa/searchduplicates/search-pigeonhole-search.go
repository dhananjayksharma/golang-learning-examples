package main

import "fmt"

func findDuplicateNum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	tortoise := arr[0]
	hare := arr[arr[0]]
	fmt.Println("tortoise, hare:", tortoise, hare)
	for tortoise != hare {
		tortoise = arr[tortoise]
		hare = arr[arr[hare]]
	}

	tortoise = 0
	for tortoise != hare {
		tortoise = arr[tortoise]
		hare = arr[hare]
	}

	return tortoise

}
func main() {
	arr := []int{5, 2, 1, 3, 5, 7, 6, 4}
	duplicate := findDuplicateNum(arr)
	fmt.Println("duplicate:", duplicate)
}
