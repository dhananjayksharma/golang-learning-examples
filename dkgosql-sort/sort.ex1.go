package main

import (
	"fmt"
	"sort"
)

func sortFunc(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func sortStrFunc(strdata []string) []string {
	sortedData := make([]string, 0, len(strdata))
	for _, v := range strdata {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedData = append(sortedData, string(s))
	}
	return sortedData
}

func sortStrSlice(strdata []string) []string {
	sort.Slice(strdata, func(i, j int) bool {
		return strdata[i] < strdata[j]
	})
	return strdata
}
func main() {
	dataint := []int{44, 1, 8, 5, 2, 7, 8, 9, 2, 1, 4}
	dataintSort := sortFunc(dataint)
	fmt.Println(dataintSort)

	datastr := []string{"tea", "bcd", "dab", "zebra"}
	datastrSort := sortStrFunc(datastr)
	fmt.Println(datastrSort)
	datastrSliceSorted := sortStrSlice(datastr)
	fmt.Println(datastrSliceSorted)
}
