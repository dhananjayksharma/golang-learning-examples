package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handler(sigchnl os.Signal, exitchnl chan int) {
	if sigchnl == syscall.SIGTERM {
		fmt.Println("Got kill signal. ")
		fmt.Println("Program will terminate now.")
		exitchnl <- 1
		// os.Exit(0)
	} else if sigchnl == syscall.SIGINT {
		fmt.Println("\nGot CTRL+C signal")
		fmt.Println("Closing.")
		exitchnl <- 1
		// os.Exit(0)
	} else {
		fmt.Println("Ignoring signal: ", sigchnl)
	}
}

func main() {
	fmt.Println("main start")
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)

	exitchnl := make(chan int, 1)

	go func() {
		for {
			s := <-sigchnl
			handler(s, exitchnl)
		}
	}()

	exitcode := <-exitchnl
	fmt.Println("exit channel:", exitcode)
}
