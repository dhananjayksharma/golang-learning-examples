package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	priceList := []int{120, 40, 60, 80, 50, 90}
	outMaxProfit := maxProfit(priceList)
	fmt.Println("max_profit:", outMaxProfit)
}

func maxProfit(prices []int) int {
	sort.Ints(prices)
	minPrice := math.MaxInt
	// fmt.Println("minPrice:", minPrice)
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
