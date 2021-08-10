package main

import (
	"fmt"
	"log"
	"net/http"
	"product-api/rest"
)

const (
	defaultIP          = ""
	defaultPort uint16 = 4000
)

func main() {

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	ip := defaultIP
	port := defaultPort
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It Works!")
	})
	s := rest.NewServer(ip, port, sm)
	log.Println("Starting Server")
	return s.Start()
}
