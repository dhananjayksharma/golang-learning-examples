package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the debug severity or above
	log.SetLevel(log.DebugLevel)

	// logrus show line number
	// log.SetReportCaller(true)

	log.Info("Info message")
	log.Warn("Warn message")
	log.Error("Error message")
	log.Fatal("Fatal message")
}
