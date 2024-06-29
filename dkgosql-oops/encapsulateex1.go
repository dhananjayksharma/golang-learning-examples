package main

import (
	sp "dkgosql-oops/shapes"
	"fmt"
)

func main() {
	// var shape sp.Shapers
	rect := sp.NewShapers()
	rect.Area(10, 5)
	fmt.Println("Area:", rect.GetValue())

}
