package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func main() {
	// Setting up log format and log level
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.DebugLevel)

	// Example log entries with different log levels
	log.Info("This is an info log.")
	log.Warn("This is a warning log.")
	log.Debug("This is a debug log.")
	log.Error("This is an error log.")

	// A simple HTTP server to simulate the service running
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Received request")
		w.Write([]byte("Hello, this is a Go service with logging!"))
	})
	
	port := ":8080"
	log.Infof("Starting server on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
