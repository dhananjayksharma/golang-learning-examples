package main

import (
	"fmt"
	"os"
)

func main() {
	a := 10
	b := 4

	a, b = b, a

	x := [5]int{1, 2, 0, 7, 3}
	fmt.Println(len(x))
	tmp(&x)
	name := "dhananjay"
	check(name)

	// cwd := os.D
	// oops -- abstraction and encap
	// godeps  ..
	// gorm
	// acid ---
	// using the function
	// gomux
	// gogin
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	// get pull
	// get fetch
}

func check(str string) {
	outmap := make(map[rune]int)
	for _, v := range str {
		if _, ok := outmap[v]; !ok {
			outmap[v] = 1
		} else {
			outmap[v] += 1
		}
	}
	fmt.Println(outmap)
}
func tmp(x *[5]int) {
	length := len(*x)
	fmt.Println("Inside len:", length)
}

// goroot && gopath
