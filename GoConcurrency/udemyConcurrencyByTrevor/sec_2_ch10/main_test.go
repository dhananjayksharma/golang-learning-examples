package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup

	wg.Add(1)
	go printSomething("round", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	status := strings.EqualFold(output, "round")
	fmt.Println("Status:", status)
	if !status {
		t.Error("Expected round but not found:", output)
	}

}
