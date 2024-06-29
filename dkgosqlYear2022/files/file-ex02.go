package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("thermopylae.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ":")
		fmt.Println(data)
		fmt.Printf("data: %#v\n", data[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
