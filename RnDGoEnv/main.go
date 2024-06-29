package main

import (
	"fmt"
)

func main() {
	fmt.Println("In go main")

	fmt.Println("Dev Configuration")
	configuration := GetConfig()
	fmt.Println("DB : ", configuration.DB_USERNAME)
}
