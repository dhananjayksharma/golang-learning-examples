package main

import "fmt"

func reverseArr(nums []int) []int {
	// outArr := make([]int, 0, len(nums)) // 0,
	numsLtr := len(nums)
	for i := (numsLtr - 1); i >= 0; i-- {
		// out := nums[i : i+1]
		// outArr = append(outArr, nums[i:i+1]...)
		nums[i-1] = int(nums[i : i+1])
	}
	fmt.Println("outArr len, cap:", len(outArr), cap(outArr))
	return nums
}

func main() {
	nums := []int{1, 2, 9, 8, 7}
	fmt.Println(nums)
	// n := nums[4:5] == 7
	outArr := reverseArr(nums)
	fmt.Println(outArr)
}
