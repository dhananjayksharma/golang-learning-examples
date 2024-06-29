package main

import "fmt"

func reverseNow(str string) {
	strLtr := len(str)
	for i := (strLtr - 1); i > 0; i-- {
		fmt.Printf("%s", string(str[i]))
	}
}

func main() {
	strdata := "golang"

	reverseNow(strdata)
}
