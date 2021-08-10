package main

import (
	"log"
	"product-api/rest"
)

const (
	defaultIP          = "localhost"
	defaultPort uint16 = 8000
)

func main() {

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ip := defaultIP
	port := defaultPort
	s := rest.NewServer(ip, port, nil)

	return s.Start()
}
