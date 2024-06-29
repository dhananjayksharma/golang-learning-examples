package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// get file from terminal
	inputFile := "mapdata.txt"
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	outMap := 0
	_ = outMap
	// convert the file binary into a string using string
	fileContent := string(file)
	// print file content
	fmt.Println(fileContent)
}
