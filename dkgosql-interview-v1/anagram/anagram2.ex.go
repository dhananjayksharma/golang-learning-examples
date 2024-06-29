package main

import (
	"fmt"
	"strings"
)

// Function to check if two strings are anagrams
func areAnagrams(str1, str2 string) bool {
	// Remove spaces and convert strings to lowercase
	str1 = strings.ReplaceAll(strings.ToLower(str1), " ", "")
	str2 = strings.ReplaceAll(strings.ToLower(str2), " ", "")

	// Check if the strings have the same length
	if len(str1) != len(str2) {
		return false
	}

	// Create a map to count the frequency of each character in str1
	charCount := make(map[rune]int)

	// Count the frequency of each character in str1
	for _, char := range str1 {
		charCount[char]++
	}

	// Compare the frequency of each character in str2
	for _, char := range str2 {
		if count, found := charCount[char]; found && count > 0 {
			charCount[char]--
		} else {
			return false
		}
	}

	// Check if all characters in str2 have been accounted for
	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test cases
	str1 := "listen"
	str2 := "silent"
	fmt.Printf("%s and %s are anagrams: %v\n", str1, str2, areAnagrams(str1, str2))

	str3 := "hello"
	str4 := "world"
	fmt.Printf("%s and %s are anagrams: %v\n", str3, str4, areAnagrams(str3, str4))
}
