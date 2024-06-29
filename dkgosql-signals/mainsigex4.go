package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Retrivie(datachan chan string, waitChan chan int) {
	for data := range datachan {
		fmt.Println("data:", data)
	}
	fmt.Println("All received")
	waitChan <- 1
}

func GenerateData(sigchan chan os.Signal, datachan chan string) {
	var s os.Signal
	go func() {
		s = <-sigchan
	}()
	for {
		time.Sleep(2 * time.Second)
		someDatetime := time.Now().GoString()
		datachan <- someDatetime
		fmt.Println("sent again ")
		if s == syscall.SIGABRT || s == syscall.SIGINT {
			fmt.Println("received os.Signal:", s)
			close(datachan)
			break
		}
	}
}

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)
	datachan := make(chan string)
	waitChan := make(chan int, 1)
	go GenerateData(sigchan, datachan)
	go Retrivie(datachan, waitChan)
	fmt.Println("Wating for go routines to terminate")
	<-waitChan
}
