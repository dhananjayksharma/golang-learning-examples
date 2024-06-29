package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {

	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))
	out := copy(destination, source)

	s.PrintBacking = true
	s.MaxPerLine = 10
	destination[2] = 9
	s.Show("source:", source)
	s.Show("destination:", destination)
	fmt.Println("out:", out)

}
