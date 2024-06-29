package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {

	var years []int

	s.PrintBacking = true
	s.MaxPerLine = 10
	s.Show("years1:", years)
	var highrainyears = make([]int, 4)
	highrainyears[1] = 2020
	s.Show("highrainyears1:", highrainyears)

	var lowrainyears []int
	s.Show("lowrainyears1:", lowrainyears)
	lowrainyears = append(lowrainyears, []int{1990, 1998}...)
	s.Show("lowrainyears2:", lowrainyears)

	lowrainyears = append(lowrainyears, []int{1991}...)
	s.Show("lowrainyears3:", lowrainyears)

	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	words = append(words[0:1], "is", "everywhere")
	s.Show("words:", words)
	words = append(words, words[len(words)+1:cap(words)]...)
	s.Show("words:", words)
}
