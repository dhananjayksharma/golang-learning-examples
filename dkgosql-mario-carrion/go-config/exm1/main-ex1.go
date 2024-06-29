package main

import (
	"flag"
	"fmt"
)

func main() {
	var password string
	flag.StringVar(&password, "password", "defaultpassword", "environment to use")
	flag.Parse()

	fmt.Printf("password: %+v\n", password)
}
