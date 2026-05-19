package main

import "fmt"


type Stack struct{
	Items []int
}


func (s *Stack) Push(val int){
	s.Items = append(s.Items, val)
}

func (s *Stack) Pop() int{
	if s.IsEmpty(){
		fmt.Println("Stack is empty")
		return 0
	}
	//val := s.Items(len(s.Items)-1)
	val := s.Items[len(s.Items)-1]
	s.Items = s.Items[0:len(s.Items)-1]
	return val
}

func (s *Stack) Peek(val int){
	
}

func (s *Stack) IsEmpty() bool{
	return len(s.Items) == 0
}

func (s *Stack) Size(val int){
	
}

func main(){

	s1 := Stack{}
	s1.Push(3)
	s1.Push(30)
	fmt.Println(s1)

}