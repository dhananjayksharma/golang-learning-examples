package main

import "fmt"

func selectionSort(nums []int) []int {
	outsort := make([]int, 0, len(nums))
	arrLen := len(nums)

	for i := 0; i < arrLen; i++ {
		minidx := i
		for j := i; j < arrLen; j++ {
			if nums[j] < nums[minidx] {
				minidx = j
			} // 11, 1, 5, 3, 6, 7, 2, 4
		}
		nums[i], nums[minidx] = nums[minidx], nums[i]
	}
	outsort = nums
	return outsort
}
func main() {
	nums := []int{11, 1, 5, 3, 6, 7, 2, 4}
	fmt.Println("Here nums:", nums)
	outsort := selectionSort(nums)
	fmt.Println(outsort)

}
