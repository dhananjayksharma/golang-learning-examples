package main

import (
	"fmt"
)

func main() {
	inputs := [][]string{
		{"eat", "tea", "ate"},
		{"me", "em"},
		{"ear", "rea", "are"},
	}
	outputList := make([][]string, 0, len(inputs))
	for _, v := range inputs {
		outlist := getPrefix(v, "_")
		outputList = append(outputList, outlist)
	}

	fmt.Println(outputList)
}
func getPrefix(list []string, pre string) []string {
	outlist := make([]string, 0, len(list))
	for _, v := range list {
		vquote := fmt.Sprintf("%q", fmt.Sprintf("%s%s", pre, v))
		outlist = append(outlist, vquote)
	}
	return outlist
}
