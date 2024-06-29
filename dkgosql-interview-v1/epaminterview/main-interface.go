package main

import "fmt"

type X2 interface {
	~int
}

type MyInt int

func (m MyInt) SomeMethod() {
	fmt.Println(m)
}

func main() {
	var x X2 = MyInt(42)
	// Use x as an instance of the X interface
	// ...
	x.SomeMethod()
	fmt.Println(x)
}
