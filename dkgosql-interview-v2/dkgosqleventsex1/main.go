package main

import (
	"fmt"
	"multithreading2/pkg/events"
	"time"
)

func main() {
	manager := events.CreateManger(time.Second / 10)

	for event := range manager.Stream() {
		fmt.Println(event.ToString())
	}
}
