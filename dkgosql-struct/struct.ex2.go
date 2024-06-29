package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJson := `{"birds":[{"pigeon":"likes to perch on rocks","eagle":"bird of prey","hawaw":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"}},{"pigeon":"likes to perch on rocks","eagle":"bird of prey"}]}`
	var result map[string]any
	err := json.Unmarshal([]byte(birdJson), &result)
	fmt.Println("Err:", err)
	fmt.Println("result:", result)
	for i, v := range result {
		fmt.Println(i, v)
	}
}
