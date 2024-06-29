package main

import (
	"fmt"
	"sort"
	"strings"
)

func ArrayChallenge(strArr []string) string {
	cacheCap := make([]string, 0, 5)
	storedData := make(map[string]int)
	// code goes here
	for i, v := range strArr {
		if _, ok := storedData[v]; !ok {
			cacheCap = append(cacheCap, v)
			storedData[v] = 1
		} else {
			if len(cacheCap) == 5 {
				cacheCap = cacheCap[1:5]
				cacheCap = append(cacheCap, v)
			} else {
				cacheCap = append(cacheCap, v)
			}

		}
		fmt.Println("i,v:", i, v)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(cacheCap)))
	strout := strings.Join(cacheCap, "-") + ":nsbrlc40"
	fmt.Println("strout:", strout)
	return strout

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	aa := []string{"A", "B", "A", "C", "A", "B"}
	fmt.Println(ArrayChallenge(aa))

}
