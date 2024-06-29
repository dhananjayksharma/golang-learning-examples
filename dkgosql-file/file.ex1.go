package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "file.txt"
	createFile(fileName)
	readFile(fileName)
}

func createFile(fileName string) {
	fmt.Printf("Writing to a file in Go lang\n")

	// in case an error is thrown it is received
	// by the err variable and Fatalf method of
	// log prints the error message and stops
	// program execution
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Defer is used for purposes of cleanup like
	// closing a running file after the file has
	// been written and main //function has
	// completed execution
	defer file.Close()

	// len variable captures the length
	// of the string written to the file.
	len, err := file.WriteString("Welcome to GeeksforGeeks." +
		" This program demonstrates reading and writing" +
		" operations to a file in Go lang.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	// Name() method returns the name of the
	// file as presented to Create() method.
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}
func readFile(fileName string) {
	fmt.Printf("\n\nReading a file in Go lang\n")

	// The ioutil package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
}
