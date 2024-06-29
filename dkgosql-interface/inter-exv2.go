package main

import "fmt"

type discount interface {
	offer() float64
}

type catalog interface {
	shipping() float64
	tax() float64
	totalAmount() float64
	discount
	displayName() string
}

type configurable struct {
	Name       string
	price, qty float64
}

type download struct {
	Name       string
	price, qty float64
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
func (c *configurable) displayName() string {
	return c.Name
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

func (d *download) tax() float64 {
	return d.price * d.qty * 0.07
}

func (d *download) shipping() float64 {
	return d.qty * 5
}
func (d *download) displayName() string {
	return d.Name
}
func display(c []catalog) {
	for _, obj := range c {
		fmt.Println("Name: ", obj.displayName())
		fmt.Println("Shipping Charge: ", obj.shipping())
		fmt.Println("Tax: ", obj.tax())
		fmt.Println("Total Amount: ", obj.totalAmount())
		fmt.Println("Total Discount: ", obj.offer())
		fmt.Println("Final Amount: ", obj.totalAmount()-obj.offer())
	}
}
func main() {
	var cat = []catalog{}
	cat = append(cat, &configurable{Name: "tshirt", price: 250, qty: 2})

	cat = append(cat, &download{Name: "Python in 24 Hours", price: 20, qty: 20})

	display(cat)
}
