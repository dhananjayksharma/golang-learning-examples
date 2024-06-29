package main

import "fmt"

func main() {
	str1 := "ear"
	str2 := "ar\U0001f600"
	str1 = "bbt"
	str2 = "cat"
	sum1 := 0
	sum2 := 0
	sum1 = sumRune(str1)
	sum2 = sumRune(str2)
	fmt.Println("is anagram sum1 : sum2 :", sum1, sum2)
	if len(str1) == len(str2) && sum1 > 0 && sum2 > 0 {
		fmt.Println("an anagram")
	} else {
		fmt.Println("not an anagram")
	}
}
func sumRune(str string) int {
	sum := 0
	for _, v := range str {
		if len(string(v)) == 1 && int(v) >= 65 && int(v) <= 122 {
			sum += int(v)
		} else {
			fmt.Println("its emmoji:", string(v), len(string(v)), v)
		}
	}
	return sum
}
