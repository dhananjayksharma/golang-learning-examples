package main

import "fmt"

func main() {
	movie := struct {
		string
		int
		bool
	}{
		"hum char",
		1995,
		true,
	}

	fmt.Println("movie:", movie)
}
