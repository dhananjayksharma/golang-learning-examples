package main

import "fmt"

func insertionSort(items []int) []int {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	return items
}
func main() {
	nums := []int{11, 1, 5, 3, 6, 7, 2, 4}
	fmt.Println("Here nums:", nums)
	outsort := insertionSort(nums)
	fmt.Println(outsort)

}
