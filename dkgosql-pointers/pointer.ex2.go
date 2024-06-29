package main

import (
	"fmt"
	"strings"
)

func main() {
	slice()
}

func slice() {
	week := []string{"wed", "mon", "tue"}
	// upSlice(week)
	fmt.Println("up-> done:", week)
	upPtrSlice(&week)
	fmt.Println("up-> done:", week)
}

func upPtrSlice(wk *[]string) {
	list := *wk
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}
	*wk = append(*wk, "thu")
}
func upSlice(wk []string) {
	// fmt.Printf("%T\n", wk)
	for i := range wk {
		wk[i] = strings.ToUpper(wk[i])
	}
}
