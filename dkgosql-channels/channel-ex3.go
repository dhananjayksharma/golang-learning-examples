package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func PDFChannel(c chan os.Signal, wg *sync.WaitGroup, orderdata chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		orderdata <- i
		time.Sleep(1 * time.Second)
	}
	close(orderdata)
	fmt.Println("PDFChannel work done!")
}

func main() {
	fmt.Println("main start")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	wtchan := make(chan struct{})
	go GetOrder(c, wtchan)
	fmt.Println("processing order in process...")
	<-wtchan
	if <-c == syscall.SIGTERM {
		fmt.Println("killing this")
	}
	fmt.Println("main end")
}

func GetOrder(c chan os.Signal, wtchan chan struct{}) {
	orderdata := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go PDFChannel(c, &wg, orderdata)
	go SendEmailChannel(c, &wg, orderdata)
	wg.Wait()
	fmt.Println("waiting for processing")
	wtchan <- struct{}{}
}
func SendEmailChannel(c chan os.Signal, wg *sync.WaitGroup, orderdata chan int) {
	defer wg.Done()
	for data := range orderdata {
		fmt.Println("reading order : orderid:", data)
	}
	fmt.Println("SendEmailChannel work done!")
}
