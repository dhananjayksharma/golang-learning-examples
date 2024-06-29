package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	t0 := time.Now()

	builder := strings.Builder{}
	for i := 0; i < 100_000; i++ {
		builder.WriteString("falcon")
	}

	t1 := time.Now()

	result := ""
	for i := 0; i < 100_000; i++ {
		result += "falcon"
	}

	t2 := time.Now()

	fmt.Println(t1.Sub(t0))
	fmt.Println(t2.Sub(t1))
}
