package main

import (
	"fmt"
	"unicode"
)

// Main function
func main() {

	// Creating rune
	rune_1 := 'g'
	rune_2 := 'e'
	rune_3 := 'E'
	rune_4 := 'k'
	rune_6 := '😀'

	// Mapping the given rune into title case
	// Using ToTitle() function
	fmt.Printf("Result 1: %c, %d ", unicode.ToTitle(rune_1), len(string(rune_1)))
	fmt.Printf("\nResult 2: %c ", unicode.ToTitle(rune_2))
	fmt.Printf("\nResult 3: %c ", unicode.ToTitle(rune_3))
	fmt.Printf("\nResult 4: %c ", unicode.ToTitle(rune_4))
	fmt.Printf("\nResult 5: %c ", unicode.ToTitle('s'))
	fmt.Printf("\nResult 6: %c ", unicode.ToTitle(rune_6))
	fmt.Printf("\nResult 6: %c, %d ", rune_6, len(string(rune_6)))

}
