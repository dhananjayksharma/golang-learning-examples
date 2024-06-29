package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	var fruits = []string{"mango", "apple", "orange", "kiwi"}

	s.PrintBacking = true
	s.MaxPerLine = 10
	s.Show("fruits1:", fruits)

	var fruitsUpper = make([]string, 0, len(fruits))

	for _, v := range fruits {
		s.Show("fruitsUpper1:", fruitsUpper)
		fruitsUpper = append(fruitsUpper, strings.ToUpper(v))
	}

}
