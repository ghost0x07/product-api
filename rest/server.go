package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Server struct {
	IP      string
	Port    uint16
	Handler http.Handler
}

func NewServer(ip string, port uint16, handler http.Handler) *Server {
	return &Server{
		IP:      ip,
		Port:    port,
		Handler: handler,
	}
}

func (s *Server) Start() error {
	port := strconv.Itoa(int(s.Port))
	addr := s.IP + ":" + port

	serv := &http.Server{
		Addr:         addr,
		Handler:      s.Handler,
		IdleTimeout:  20 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	go func() {
		err := serv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	log.Println("Received terminate signal:", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := serv.Shutdown(tc)
	if err != nil {
		return fmt.Errorf("unable to gracefully shutdown server: %w", err)
	}
	return nil
}
