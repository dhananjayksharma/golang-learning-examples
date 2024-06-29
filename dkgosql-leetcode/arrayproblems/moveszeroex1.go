package main

import "fmt"

func moveZeroes(nums []int) {
	out := make([]int, len(nums))
	zero := len(nums) - 1
	nonezero := 0

	for _, v := range nums {
		if v == 0 {
			out[zero] = 0
			zero--
		} else {
			out[nonezero] = v
			nonezero++
		}
	}
	fmt.Println(out)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
