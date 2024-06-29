// https://play.golang.org/p/PFUHW5d-26u
package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Result []int `json:"result"`
}

func main() {
	// nil slice
	var s1 []int
	r1 := &Response{
		Result: s1,
	}
	s1 = append(s1, 11)
	fmt.Printf("s1 %+v\n", s1)

	b, _ := json.Marshal(r1)
	fmt.Printf("%+v\n", string(b)) // {"result":null}

	s2 := []int{}
	s2 = append(s2, 88)
	fmt.Printf("s2 %+v\n", s2)

	r2 := &Response{
		Result: s2,
	}

	b, _ = json.Marshal(r2)
	fmt.Printf("%+v\n", string(b)) // {"result":[]}
}
