package main

import "fmt"

func bubbleSort(nums []int) []int {
	// outsort := []int{}
	arrLen := len(nums)
	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}
func main() {
	nums := []int{11, 1, 5, 3, 6, 7, 2, 4}
	fmt.Println("Here nums:", nums)
	outsort := bubbleSort(nums)
	fmt.Println(outsort)

}
