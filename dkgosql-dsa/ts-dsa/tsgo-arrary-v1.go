package main

import "fmt"

func arrayDemo() {
	// Fixed-size array
	var arr [5]int = [5]int{10, 20, 30, 40, 50}

	// Slice (dynamic array) — preferred in Go
	nums := []int{1, 2, 3, 4, 5}

	// 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Fixed-size array:", arr)
	fmt.Println("Slice:", nums)
	fmt.Println("2D slice:", matrix)
}

func reverseInPlaceVer1(ids []int) {
	l, r := 0, len(ids)-1
	for l < r {
		ids[l], ids[r] = ids[r], ids[l]
		l++
		r--
	}
}

// Reverse array in-place — O(n) time, O(1) space
func reverseInPlace(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func reverseInPlaceVer2(ids []int) {
	l, r := 0, len(ids)-1
	for l < r {
		ids[l], ids[r] = ids[r], ids[l]
		l++
		r--
	}
}

// Sliding window — max sum subarray of size k
func maxSumSubarray(arr []int, k int) int {
	windowSum, maxSum := 0, 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum = windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func main() {
	arrayDemo()
	arr := []int{1, 9, 5, 7, 15, 10, 25}
	out := maxSumSubarray(arr, 3)
	fmt.Println("out maxSumSubarray:", out)

	// Slice (dynamic array) — preferred in Go
	nums := []int{1, 2, 6, 3, 4, 5}
	// Two-pointer technique
	reverseInPlace(nums)
	fmt.Println("Slice reverseInPlace:", nums)

	ids := []int{9, 5, 2, 4, 6, 15, 21, 11}
	// Two-pointer technique
	reverseInPlaceVer1(ids)
	fmt.Println("Slice reverseInPlaceVer1:", ids)

    groups := []int{80, 55, 1, 12, 6, 35 , 4}
	// Two-pointer technique
	reverseInPlaceVer2(groups)
	fmt.Println("Slice reverseInPlaceVer2:", groups)
}
