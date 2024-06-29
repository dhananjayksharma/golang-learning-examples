package main

import "fmt"

//Maximum distance between two occurrences of same element in array

// maximum distance for 2 is 11-1 = 10
// maximum distance for 1 is 4-2 = 2
// maximum distance for 4 is 10-5 = 5
type position struct {
	start int
	next  int
	max   int
}

func findMaxDistance(nums []int) (int, int) {
	ele, out := 0, 0
	hashMap := make(map[int]position)
	for i, v := range nums {
		if _, ok := hashMap[v]; !ok {
			hashMap[v] = position{start: i, next: 0, max: 0}
		} else {
			data := hashMap[v] // s:1, m:0, n:0 --
			hashMap[v] = position{start: data.start, next: i, max: i - data.start}
			if hashMap[v].max > out {
				out = hashMap[v].max
				ele = v
			}
		}
	}
	fmt.Printf("hashMap:%+v\n", hashMap)
	return ele, out
}
func main() {
	input := []int{3, 2, 1, 2, 1, 4, 5, 8, 6, 7, 4, 2} //, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
	//2 - 1, 2-11 = 10, 4-5, 4-10 = 5
	// input := []int{3, 2, 1}
	// Output: 10
	ele, out := findMaxDistance(input)
	fmt.Println("Element, out: ", ele, out)
}
