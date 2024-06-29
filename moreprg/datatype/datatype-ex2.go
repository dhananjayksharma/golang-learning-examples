package main

import "fmt"

func main(){
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Printf("%T", p)
}