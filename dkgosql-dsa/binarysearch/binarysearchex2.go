package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		median := (low + high) / 2
		fmt.Println("median:", median)
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	fmt.Println("low:", low)
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	out := binarySearch(119, items)
	fmt.Println("Out: ", out)
}
