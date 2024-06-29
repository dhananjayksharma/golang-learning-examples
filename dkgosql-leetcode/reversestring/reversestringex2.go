package main

import (
	"fmt"
	"strings"
)

func Reverse(inputstr string, outstr *string) string {
	builder := strings.Builder{}
	for i := len(inputstr) - 1; i >= 0; i-- {
		_, _ = builder.WriteString(string(inputstr[i]))
	}
	return builder.String()
}
func main() {
	sentence := "I am a golang developer!"
	outstr := ""
	out := Reverse(sentence, &outstr)
	fmt.Println(out)
}
