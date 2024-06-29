package main

import "fmt"

//DemoStruct definition
type DemoStruct struct {
    Val int
}
//A.
func demo_func1() DemoStruct {
    return DemoStruct{Val: 1}
}
//B.
func demo_func2() *DemoStruct {
    return &DemoStruct{}
}
//C.
func demo_func3(s *DemoStruct) {
    s.Val = 1
}

func main(){
	fmt.Printf("func 1: %v\n", demo_func1())

	fmt.Printf("func 2: %v\n", demo_func2())
	s :=DemoStruct{}
	demo_func3(&s)
	fmt.Printf("func 3: %v\n", s.Val)
}