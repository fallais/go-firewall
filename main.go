package main

import (
	"flag"
	"net"
	"net/http"
	"time"

	"go-firewall/connectors/cisco"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/zenazn/goji/graceful"
)

var (
	bindAddress = flag.String("bind", ":8000", "Network address used to bind")
	logging     = flag.String("logging", "info", "Logging level")
)

func init() {
	// Parse the flags
	flag.Parse()

	// Set localtime to UTC
	time.Local = time.UTC

	// Set the logging level
	level, err := logrus.ParseLevel(*logging)
	if err != nil {
		logrus.Fatalln("Invalid log level ! (panic, fatal, error, warn, info, debug)")
	}
	logrus.SetLevel(level)

	// Set the TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	logrus.Infoln("go-firewall is starting")
}

func main() {
	c, err := cisco.NewCiscoASA("lab.sharontools.com:22", "titi", "toto")
	if err != nil {
		logrus.Fatalln(err)
	}
	_, err = c.GetConfiguration()
	if err != nil {
		logrus.Fatalln(err)
	}

	// Merge the routes
	http.Handle("/metrics", prometheus.Handler())

	// Starting the HTTP server
	logrus.Infoln("Starting the HTTP server on interface", *bindAddress)

	// Create the listener
	listener, err := net.Listen("tcp", *bindAddress)
	if err != nil {
		logrus.Fatalln("Cannot set up a TCP listener :", err)
	}

	// Handle the signals with graceful
	graceful.HandleSignals()

	// Start the listening
	err = graceful.Serve(listener, http.DefaultServeMux)
	if err != nil {
		logrus.Error(err)
	}

	// Wait until open connections close
	graceful.Wait()
}
