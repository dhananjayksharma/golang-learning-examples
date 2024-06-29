package main

import "fmt"

func main() {
	var list = []int{1, 2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	var target int = 91
	in, out := search(target, list)
	fmt.Println("Found in :", in, out)
}

func search(target int, source []int) (string, int) {
	len := len(source) / 2

	if source[len] == target {
		return "middle", source[len]
	}
	leftSource := source[:len]
	fmt.Println("leftSource:", leftSource)
	leftOut := checkTarget(target, leftSource)
	if target == leftOut {
		return "left", leftOut
	}

	rightSource := source[len:]
	fmt.Println("rightSource:", rightSource)
	rightOut := checkTarget(target, rightSource)
	if target == rightOut {
		return "right", rightOut
	}

	return "", 0
}
func checkTarget(target int, source []int) int {
	for _, v := range source {
		if v == target {
			return v
		}
	}
	return 0
}
