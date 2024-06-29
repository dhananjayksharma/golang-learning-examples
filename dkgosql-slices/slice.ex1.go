package main

import "fmt"

func main() {
	// slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// fmt.Println(len(slice), cap(slice)) //Print 5, //Print 5
	// slice = slice[4:]
	// fmt.Println(slice, len(slice), cap(slice)) //Print 5, //Print 5
	intSlice := []int{10, 60, 70}
	fmt.Println(len(intSlice), cap(intSlice)) //Print 3 //Print 3
	intSlice = append(intSlice, intSlice[1:3]...)
	intSlice = append(intSlice[:1], 30, 40)

	fmt.Println(len(intSlice), cap(intSlice), intSlice[:5]) //Print 3 //Print 6
}
