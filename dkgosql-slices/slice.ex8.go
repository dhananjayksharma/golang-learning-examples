package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	var years = []int{1, 2, 3, 4}

	s.PrintBacking = true
	s.MaxPerLine = 10
	s.Show("years1:", years)

	// version ::: 0
	// lowrainyears := years[:2:2]
	// s.Show("lowrainyears1:", lowrainyears)

	// // lowrainyears = append(lowrainyears, 10, 20)
	// s.Show("lowrainyears1:", lowrainyears)

	// s.Show("years2:", years)

	// version ::: 1
	lowrainyears := append(years[:2:2], 10, 20)
	s.Show("lowrainyears1:", lowrainyears)
	s.Show("years2:", years)
}
