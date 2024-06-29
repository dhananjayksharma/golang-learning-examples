package main

import (
	"fmt"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello World!"

	wg.Add(1)
	go updateMessage("x")
	wg.Wait()
	wg.Add(1)
	go updateMessage("Hello Universe!")
	wg.Wait()

	fmt.Println("msg :", msg)
	if msg != "Hello Universe!" {
		t.Error("incorrect value in msg")
	}
}
