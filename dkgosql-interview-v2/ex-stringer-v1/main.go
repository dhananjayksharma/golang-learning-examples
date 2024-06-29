package main

import "fmt"

type Rectangle struct {
	Width int
	Legth int
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Length:%d, Width:%d", r.Legth, r.Width)
}

func main() {
	rc := Rectangle{Legth: 12, Width: 10}
	fmt.Println(&rc)
	fmt.Printf("%v", &rc)
}
