package main

import (
	"fmt"
	"log"
	"product-api/rest"
)

var host = ""
var port uint16 = 4000

func main() {

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	err := parseEnv()
	if err != nil {
		return fmt.Errorf("unable to parse env variables: %w", err)
	}

	r := createRouter()

	s := rest.NewServer(host, port, r)
	log.Printf("Starting Server on %v:%v\n", host, port)

	return s.Start()
}
