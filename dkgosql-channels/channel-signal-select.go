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
	for i := 1; i <= 10; i++ {
		uniqueTime := time.Now()
		uuid <- time.Time.String(uniqueTime) + " sr #" + fmt.Sprintf("%d", i)
	}
	close(uuid)
}

func PrintUUID(uuid chan string, waitChannel chan struct{}, sigs chan os.Signal) {
	for v := range uuid {
		time.Sleep(1 * time.Second)
		fmt.Println("Number:", v)
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
	fmt.Println("end")
}
