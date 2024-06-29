package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"a":"apple","key":"val1","key3":"val2"}`
	x := map[string]string{}

	err := json.Unmarshal([]byte(jsonStr), &x)
	fmt.Printf("%v", err)
	fmt.Println(x["key3"])
}
