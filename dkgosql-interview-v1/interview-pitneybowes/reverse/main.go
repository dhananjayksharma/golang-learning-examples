package main

import (
	"fmt"
	"strings"
)

func getReverse(str string) string {
	var outstr string
	strLen := len(str) - 1
	for i := strLen; i >= 0; i-- {
		outstr = outstr + string(str[i])
	}
	return outstr
}

func getReverseRune(str string) (string, []string) {
	var outstr string
	var outstrslice = []string{}
	for _, v := range str {
		outstr = outstr + string(v)
		outstrslice = append(outstrslice, string(v))
	}
	outSlice := make([]string, 0, len(outstrslice))
	for i, _ := range outstrslice {
		start := len(outstrslice) - 1 - i
		outSlice = append(outSlice, outstrslice[start:start+1]...)
	}
	// fmt.Println("len(outstrslice):", len(outstrslice), outSlice)
	finalStr := strings.Join(outSlice, "")
	return finalStr, outSlice
}

func main() {
	fmt.Println("start")
	inputstr := "golan😀gi😀seasy" // str = [0:1]
	outstr := getReverse(inputstr)
	fmt.Printf("input:%s,\noutput:%s\n", inputstr, outstr)

	outstr, outstrslice := getReverseRune(inputstr)
	fmt.Printf("input:%s,\noutput:%s\n slice:%v\n", inputstr, outstr, outstrslice)
}
