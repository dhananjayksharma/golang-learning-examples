package main

import (
	"fmt"
)

func findMaxSubarraySum(nums []int) int {
	maxSum := nums[0]     // initialize maxSum with the first element
	currentSum := nums[0] // initialize currentSum with the first element
	fmt.Println("nums:", nums)
	for i := 1; i < len(nums); i++ {
		// Check if adding the current element increases the sum or if the current element alone is greater
		fmt.Println("nums[i], currentSum, currentSum+nums[i], :", nums[i], currentSum, currentSum+nums[i])
		currentSum = max(nums[i], currentSum+nums[i])
		// Update maxSum if the currentSum is greater

		maxSum = max(maxSum, currentSum)

		fmt.Println(" maxSum, currentSum:", maxSum, currentSum)
	}

	return maxSum
}

// Helper function to find the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// nums := []int{2, 3, -1, 4, -10, 2, 5}
	// maxSum := findMaxSubarraySum(nums)
	// fmt.Println("Max subarray sum:", maxSum) // Output: 8

	// nums = []int{2, 3, -6, 4, 2, -8, 3}
	// maxSum = findMaxSubarraySum(nums)
	// fmt.Println("Max subarray sum:", maxSum) // Output: 6

	nums := []int{-1, -2, 4, -3, -5}
	// nums = []int{2, 4, 3, 5, 2}
	maxSum := findMaxSubarraySum(nums)
	fmt.Println("Max subarray sum:", maxSum) // Output: 6
}
