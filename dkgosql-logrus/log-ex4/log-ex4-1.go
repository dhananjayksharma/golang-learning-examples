package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	logFile := "log.txt"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		panic(err)
	}
	defer f.Close()
	// Output to stdout instead of the default stderr
	log.SetOutput(f)

	// Only log the debug severity or above
	log.SetLevel(log.DebugLevel)

	log.Info("Info message")
	log.Warn("Warn message")
	log.Error("Error message")
	log.Fatal("Fatal message")
}
