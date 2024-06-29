package main

import (
	"fmt"
	"sort"
)

func sortString(str string) string {
	s := []byte(str)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println("string(s):", string(s))
	return string(s)
}

func getAngramGroup(inputs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range inputs {
		sortedStr := sortString(str)
		fmt.Println("sortedStr:", sortedStr)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
		fmt.Println("output:", anagramMap)
	}

	groupedAnagrams := make([][]string, 0, len(anagramMap))
	for _, anagrams := range anagramMap {
		groupedAnagrams = append(groupedAnagrams, anagrams)
	}

	// fmt.Println("groupedAnagrams:", groupedAnagrams)
	return groupedAnagrams
}
func main() {
	inputs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "teae", "teaa", "ear", "are"}
	out := getAngramGroup(inputs)
	fmt.Println("out:", out)
}
