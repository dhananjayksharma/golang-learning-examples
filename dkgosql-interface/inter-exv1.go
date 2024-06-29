package main

import "fmt"

type discount interface {
	offer() float64
}

type catalog interface {
	shipping() float64
	tax() float64
	totalAmount() float64
}

type configurable struct {
	name       string
	price, qty float64
}

type download struct {
	name       string
	price, qty float64
}

func (d *download) tax() float64 {
	return d.price * d.qty * 0.07
}

func (c *configurable) tax() float64 {
	return c.price * c.qty * 0.05
}

func (c *configurable) shipping() float64 {
	return c.qty * 5
}

func (c *configurable) totalAmount() float64 {
	return c.price * c.qty
}

func (c *configurable) offer() float64 {
	return c.price * c.qty * 0.3
}

func (d *download) offer() float64 {
	return d.price * d.qty * 0.15
}

func (d *download) totalAmount() float64 {
	return d.price * d.qty
}

func main() {
	obj := configurable{name: "tshirt"}
	obj.price = 250
	obj.qty = 2
	fmt.Println("Name: ", obj.name)
	fmt.Println("Shipping Charge: ", obj.shipping())
	fmt.Println("Tax: ", obj.tax())
	fmt.Println("Total Amount: ", obj.totalAmount())
	fmt.Println("Total Discount: ", obj.offer())
	fmt.Println("Final Amount: ", obj.totalAmount()-obj.offer())

	book := download{"Python in 24 Hours", 20, 20}
	fmt.Println("\nDownloadable Product")
	fmt.Println("Tax: ", book.tax())
	fmt.Println("Total Amount: ", book.totalAmount())
	fmt.Println("Total Discount: ", book.offer())
	fmt.Println("Final Amount: ", book.totalAmount()-book.offer())
}
