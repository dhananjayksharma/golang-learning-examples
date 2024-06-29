package main

import "fmt"

func merge(x []int, y []int) []int {
	nums := make([]int, 0, len(x)+len(y))

	nums = append(nums, x...)
	nums = append(nums, y...)

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
	x := []int{1, 2, 4}
	y := []int{1, 3, 4}
	o := merge(x, y)
	fmt.Println(o)
}
