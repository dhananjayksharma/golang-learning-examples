package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("start")
	fmt.Println("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running darwin:", os)
	case "linux":
		fmt.Println("Running linux:", os)
	default:
		fmt.Println("Running default:", os)
	}
}
