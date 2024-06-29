package main

import (
	"fmt"
	"log"

	"dkgosql.com/match-prefix-service/internal/matchprefixes"
)

// @title Prefix finder
// @version 1.0.0
// @contact Dhananjayksharma@gmail.com
// @BasePath /truecaller
func main() {
	fmt.Println("True caller start")
	inputStr := "xyi" //"xyiuoXX" //xyi

	var pf matchprefixes.PrefixFinder

	var ps matchprefixes.Prefix
	pf = &ps
	err := pf.FindPrefix(inputStr)
	if err != nil {
		log.Printf("Error found while calling matchprefixes.FindPrefix(inputStr) Err: %v", err)
	} else {
		fmt.Println("Final Prefixes List: ", ps)
	}
}
