package main

import (
	"fmt"
	"sort"
)

func main() {
	priceList := []int{120, 40, 60, 80, 50, 90, 40}
	duplicate := containsDuplicate(priceList)

	fmt.Println("duplicate:", duplicate)
	sortedduplicate := containsDuplicateSorted(priceList)
	fmt.Println("sorted duplicate:", sortedduplicate)
}

func containsDuplicate(prices []int) int {
	seen := make(map[int]struct{})
	for _, price := range prices {
		if _, ok := seen[price]; ok {
			return price
		} else {
			seen[price] = struct{}{}
		}
	}
	return 0
}

func containsDuplicateSorted(prices []int) bool {
	sort.Ints(prices)
	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1] {
			return true
		}
	}
	return false
}
