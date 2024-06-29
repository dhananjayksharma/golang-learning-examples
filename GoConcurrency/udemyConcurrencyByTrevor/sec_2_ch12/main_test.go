package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	msg = "testround"
	go updateMessage(msg)
	wg.Wait()
	if msg != "testround" {
		t.Error("Expected value is \"testround\"", msg)
	}
}
func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	msg = "Hello, universe!"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	fmt.Println("msg :", output)
	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected value is \"Hello, universe!\"", output)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	main()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected value is \"Hello, universe!\"", output)
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected value is \"Hello, cosmos!\"", output)
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected value is \"Hello, world!\"", output)
	}
}
