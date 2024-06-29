package main

import "fmt"

type EmployeeType string

const (
	PERMANENT  EmployeeType = "PERMANENT"
	CONSULTANT EmployeeType = "CONSULTANT"
)

type Employeerer interface {
	Add()
	Count()
}

type Permanent struct {
	Tax      float64
	Name     string
	Employer Company
}

type Consultant struct {
	Name     string
	Tds      float64
	Employer Company
}
type Company struct {
	Name string
	Type string
}

func (p *Permanent) Add() {

}
func (p *Permanent) Count() {

}
func (p *Permanent) CalculatePF() {
	fmt.Println("PF:", p.Tax*.14)
}

func (c *Consultant) Add() {

}
func (c *Consultant) Count() {

}

func main() {
	// var emp Employeerer
	var emp Employeerer = &Permanent{Name: "Dhananjay", Tax: 30.0, Employer: Company{Name: "PaceNow", Type: string(PERMANENT)}}
	// fmt.Printf("Employee is %v \n", emp)
	// emp.CalculatePF()

	Display(emp)
	emp = &Consultant{Name: "Dhananjay", Tds: 10.2, Employer: Company{Name: "BundleN", Type: string(CONSULTANT)}}
	// // fmt.Printf("Employee is %v \n", emp)
	Display(emp)

	// pf, ok := emp.(*Permanent)
	// if ok == true {
	// 	pf.CalculatePF()
	// }

}

func Display(emp Employeerer) {
	switch typeEmp := emp.(type) {
	case *Permanent:
		fmt.Println("Its Employee:", typeEmp)
		typeEmp.CalculatePF()
	case *Consultant:
		fmt.Println("Its Consultant:", typeEmp)
	default:
		fmt.Println("Its default:", typeEmp)
	}
}
