package main

import s "github.com/inancgumus/prettyslice"

func main() {

	grades := []int{20, 30, 40, 50, 60, 70}

	s.PrintBacking = true
	s.MaxPerLine = 10
	grades = grades[0:0]
	s.Show("grades1:", grades)
	grades = grades[:cap(grades)]
	s.Show("grades2:", grades)

	grades = grades[2:]
	s.Show("grades3:", grades)

	grades = grades[1:]
	s.Show("grades4:", grades)
	grades = grades[1:3]
	s.Show("grades5:", grades)

	grades = grades[cap(grades):cap(grades)]
	s.Show("grades6:", grades)

	grades = nil
	s.Show("grades7:", grades)
}
