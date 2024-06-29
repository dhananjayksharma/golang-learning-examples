package main

import "fmt"

type Employeer interface {
	Add(empList *[]Employee)
}

type Employee struct {
	Name string
	Age  int32
}

func (em Employee) Add(empList *[]Employee) {
	emNew := Employee{Age: em.Age, Name: em.Name}
	*empList = append(*empList, emNew)
}

func main() {
	empList := []Employee{}
	var em Employeer

	emp1 := Employee{"Dhananjay", 35}
	em = emp1
	em.Add(&empList)
	fmt.Println("Employee list:", empList)

	emp2 := Employee{"Anupama", 16}
	em = emp2
	em.Add(&empList)
	fmt.Println("Employee list:", empList)
}
