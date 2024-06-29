package main

import "fmt"

// Skipping Values
const (
	_ int = iota
	RESP
	PM
)

func main() {

	fmt.Printf("RESP:%v, PM:%v\n", RESP, PM)

}
