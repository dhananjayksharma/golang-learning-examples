package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl, syscall.SIGINT, syscall.SIGKILL)

	waitChan := make(chan bool, 1)
	fmt.Println("main start")
	go func() {
		s := <-sigchnl
		handle(s, waitChan)
	}()
	fmt.Println("Waiting for waitChan")
	fmt.Println("main done :", <-waitChan)
	fmt.Println("main end")
}

func handle(s os.Signal, waitChan chan bool) {
	switch s {
	case syscall.SIGINT:
		fmt.Println("SignalINT", s)
		waitChan <- true
		// os.Exit(0)
	default:
		fmt.Println("Unhandled/unknown signal")
	}
}
