package main

import "fmt"

func bubbleSort(nums []int) []int {
	outsort := []int{}
	isswaprequired := false
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] > nums[i+1] {
			isswaprequired = true
			tmp := nums[i]
			nums[i] = nums[i+1]
			nums[i+1] = tmp
		}
	}
	outsort = nums
	if isswaprequired {
		fmt.Println("Here:", outsort)
		bubbleSort(outsort)
	}
	return outsort
}
func main() {
	nums := []int{11, 1, 5, 3, 6, 7, 2, 4}
	fmt.Println("Here nums:", nums)
	outsort := bubbleSort(nums)
	fmt.Println(outsort)

}
