package main

import "fmt"

func GetTwoSum(nums []int, target int) []int {
	twoSumMap := make(map[int]int, len(nums)) // , len(nums)
	// twoSumSlice := make([]int, 0, len(nums))

	for i, v := range nums {
		// twoSumSlice = append(twoSumSlice)
		// fmt.Printf("len:%d, cap:%d, twoSumSlice:%v\n\n", len(twoSumSlice), cap(twoSumSlice), twoSumSlice)
		if j, ok := twoSumMap[target-v]; ok {
			fmt.Printf("j:%d,v:%d\n", j, i)
			twoSumMap[v] = i

			return []int{j, i}
		}
		twoSumMap[v] = i
		// twoSumSlice = append(twoSumSlice, i)
	}
	fmt.Printf("twoSumMap:%#v\n", twoSumMap)
	return nil
}

func main() {
	nums := []int{7, 6, 4, 9, 11, 3, 5}

	fmt.Printf("nums:%v\n", nums)

	out := GetTwoSum(nums, 11)

	fmt.Printf("output:%v\n", out)
}
