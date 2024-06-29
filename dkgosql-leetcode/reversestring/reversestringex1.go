package main

import "fmt"

func Reverse(inputstr string, outstr *string) string {
	out := ""

	for i := len(inputstr) - 1; i >= 0; i-- {
		out = out + string(inputstr[i])
	}
	return out
}
func main() {
	sentence := "I am a golang developer!"
	outstr := ""
	out := Reverse(sentence, &outstr)
	fmt.Println(out)
}
