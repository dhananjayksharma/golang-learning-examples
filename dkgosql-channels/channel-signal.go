// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GenerateUUID(uuid chan string) {
	i := 1
	for {
		uniqueTime := time.Now()
		uuid <- time.Time.String(uniqueTime) + " sr #" + fmt.Sprintf("%d", i)
		i++
	}
}

func PrintUUID(uuid chan string, waitChannel chan struct{}, sigs chan os.Signal) {

	for i := 0; i < 25; i++ {
		fmt.Println("Number:", <-uuid)
		time.Sleep(1 * time.Second)
	}
	sig := <-sigs
	fmt.Println("Here sig:", sig)
	waitChannel <- struct{}{}
	close(waitChannel)
}

func main() {
	fmt.Println("main start")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	waitChannel := make(chan struct{})
	uuid := make(chan string)

	go GenerateUUID(uuid)
	go PrintUUID(uuid, waitChannel, sigs)
	fmt.Println("awaiting signal")
	fmt.Printf("all done %v\n", <-waitChannel)
	close(sigs)
	close(uuid)
	fmt.Println("end")
}
