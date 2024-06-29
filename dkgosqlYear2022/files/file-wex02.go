package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"sky", "falcon", "rock", "hawk"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}
