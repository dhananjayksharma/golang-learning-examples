package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetLevel(log.TraceLevel)
	// logrus.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
	log.Traceln("Trace Level")
	log.Debugln("Debug Level")
	log.Infoln("Info Level")
	log.Warningln("Warning Level")
	log.Errorln("Error Level")
	log.Fatalln("Fatal Level")
	log.Panicln("Panic Level")
}
