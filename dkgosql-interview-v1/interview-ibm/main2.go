package main

import (
	"fmt"
	"sort"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func getAngramGroup(inputs []string) [][]string {
	output := make(map[string][]string) // ear == are == after aer
	//["aer":[""]]

	lenInput := len(inputs) - 1
	for i := 0; i < lenInput; i++ {
		str1 := inputs[i]
		for j := i; j < lenInput; j++ {
			str2 := inputs[j+1]
			if len(str1) == len(str2) {
				s1 := sortString(str1)
				status := isAnagram(str1, str2)
				if status {
					if ok := contains(output[s1], fmt.Sprintf("%q", str1)); !ok {
						output[s1] = append(output[s1], fmt.Sprintf("%q", str1))
					}
					if ok := contains(output[s1], fmt.Sprintf("%q", str2)); !ok {
						output[s1] = append(output[s1], fmt.Sprintf("%q", str2))
					}
				} else {
					if ok := contains(output[s1], fmt.Sprintf("%q", str1)); !ok {
						output[s1] = append(output[s1], fmt.Sprintf("%q", str1))
					}
					s2 := sortString(str2)
					if ok := contains(output[s2], fmt.Sprintf("%q", str2)); !ok {
						output[s2] = append(output[s2], fmt.Sprintf("%q", str2))
					}
				}
			}
		}
	}

	groupedAnagrams := make([][]string, 0, len(output))
	for _, anagrams := range output {
		groupedAnagrams = append(groupedAnagrams, anagrams)
	}
	return groupedAnagrams
}

func sortString(str string) string {
	s := []byte(str)
	sort.Slice(s, func(i, j int) bool {
		return str[i] < str[j]
	})
	return string(s)
}
func isAnagram(str1, str2 string) bool {
	var runeSum = make(map[rune]int)

	for _, v := range str1 {
		if _, ok := runeSum[v]; !ok {
			runeSum[v] = 1
		} else {
			runeSum[v] += 1
		}
	}
	for _, v := range str2 {
		if _, ok := runeSum[v]; ok {
			runeSum[v] -= 1
		}
	}

	for _, v := range runeSum {
		if v > 0 {
			return false
		}
	}

	return true
}
func main() {
	inputs := []string{"eat", "tea", "ear", "are", "cat"} //,"tan", "ate", "nat", "bat", "teae", "teaa"
	out := getAngramGroup(inputs)
	fmt.Println(out)
}

// are == ear   --- called anagram
