package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	if os.Getenv("ENV") == "production" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
		// Only log the warning severity or above.
		log.SetLevel(log.DebugLevel)
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
}
