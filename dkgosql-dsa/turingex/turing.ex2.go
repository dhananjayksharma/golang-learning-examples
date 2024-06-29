package main

import "fmt"

func findCount(a, b string) int {
	count := 0
	if len(a) == 0 || len(b) == 0 {
		return count
	}
	mapCount := make(map[rune]int)
	fmt.Println("mapCount:", mapCount)
	for _, v := range b {
		if _, ok := mapCount[v]; !ok {
			mapCount[v] = 1
		} else {
			mapCount[v] += 1
		}
	}
	for _, v := range a {
		if _, ok := mapCount[v]; ok {
			mapCount[v] += 1
		}
		if count < mapCount[v] {
			count = mapCount[v]
		}
	}
	return count
}
func main() {
	a := "abc"
	b := "aabbbb"
	o := findCount(a, b)
	fmt.Println(o)
}
