package main

import (
	"fmt"
)

func firstcharacterrepeate(str string) string {
	output := make(map[rune]int)
	for _, s := range str {
		fmt.Println(output)
		if _, ok := output[s]; ok {
			return string(s)
		} else {
			output[s] = 1
		}
	}
	return ""
}
func main() {
	input := "happy"
	out := firstcharacterrepeate(input)
	if len(out) == 1 {
		fmt.Println("Found:", out)
	} else {
		fmt.Println("Not Found:", out)
	}
}
