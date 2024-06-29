package main

import "fmt"

// increase from 0 to
const (
	ROLE int = iota
	RESP
	PM
)

func main() {

	fmt.Printf("ROLE:%v, RESP:%v, PM:%v\n", ROLE, RESP, PM)

}
