package main

import "fmt"

func GetTwoSum(nums []int, target int) []int {
	twoSumMap := make(map[int]int, len(nums)) // , len(nums)
	// twoSumSlice := make([]int, 0, len(nums))

	for i, v := range nums {
		// twoSumSlice = append(twoSumSlice)
		// fmt.Printf("len:%d, cap:%d, twoSumSlice:%v\n\n", len(twoSumSlice), cap(twoSumSlice), twoSumSlice)
		fmt.Println("twoSumMap:", twoSumMap, "\n")
		if j, ok := twoSumMap[target-v]; ok {
			fmt.Printf("j value:%d, i index:%d\n", j, i)
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

	out := GetTwoSum(nums, 20)

	if len(out) == 2 {
		fmt.Printf("output:%v, v1:%d, v2:%d\n", out, nums[out[0]], nums[out[1]])
	} else {
		fmt.Println("Target not found")
	}

}
