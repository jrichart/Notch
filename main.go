package main

import (
	"Notch/server"
	"fmt"
	"log"
	"os"
)

var service string

func init() {
	log.SetOutput(os.Stdout)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname:port\n", os.Args[0])
		os.Exit(1)
	}
	service = os.Args[1]
}

func main() {
	log.Println("Starting")

	server.Run(service)
}
