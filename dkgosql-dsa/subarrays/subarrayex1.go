package main

import "fmt"

func FindSubArray(arr []int) int {
	// subarrays := make([][]int, len(arr))
	maxSum := arr[0]
	actualsum := 0
	// fmt.Println("maxSum:", maxSum)
	// fmt.Println("subarrays:len, cap, subarrays:", len(subarrays), cap(subarrays), subarrays)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			actualsum = 0
			for k := i; k <= j; k++ {
				actualsum += arr[k]
				maxSum = getMax(maxSum, actualsum)
			}
		}
	}
	return maxSum
	// fmt.Println("subarrays:len, cap, subarrays:", len(subarrays), cap(subarrays), subarrays)
}

// -1
// -2
// -3
// -1 + -2
// -2 + -3

func getMax(maxval, newval int) int {
	if newval > maxval {
		return newval
	}
	return maxval
}

func main() {
	var arr = []int{4, -7, 1, 5}
	arr = []int{2, 3, -1, 4, -10, 2, 5}
	arr = []int{2, 4, 3, 2, 5, 2}
	maxSum := FindSubArray(arr)
	fmt.Println("maxSum:", maxSum)

}
