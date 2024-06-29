package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	// grades := [...]int{20, 30, 40, 50, 60, 70}
	grades := []int{20, 30, 40, 50, 60, 70}
	front := grades[:3]
	s.PrintBacking = true
	s.MaxPerLine = 10
	s.Show("grades:", grades[:])
	s.Show("front:", front)

	// fmt.Printf("grades:%T\n", grades)
	// newgrades := grades[:]
	// fmt.Printf("newgrades:%T\n", newgrades)
}
