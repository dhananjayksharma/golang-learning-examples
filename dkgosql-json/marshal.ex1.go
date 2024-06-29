package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string
	Age      int16
	Address  any
	Address2 interface{}
}

func main() {
	var p = Person{Name: "Dhananjay", Age: 34, Address: nil, Address2: nil}
	pbyte, _ := json.Marshal(p)
	var p2 Person
	_ = json.Unmarshal(pbyte, &p2)
	fmt.Printf("%T, %+v, %v\n", p2, p2, p2)

	address := map[string]any{"plot": 112, "pin": 800115, "street": "FCI"}
	var p3 = Person{Name: "Dhananjay", Age: 34, Address: address, Address2: address}

	pbyte3, _ := json.Marshal(p3)
	var p4 Person
	_ = json.Unmarshal(pbyte3, &p4)
	fmt.Printf("%T, %+v, %v\n", p4, p4, p4)
}
