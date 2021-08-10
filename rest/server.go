package rest

import (
	"fmt"
	"net/http"
	"strconv"
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
	var err error
	port := strconv.Itoa(int(s.Port))
	addr := s.IP + ":" + port

	err = http.ListenAndServe(addr, s.Handler)
	if err != nil {
		return fmt.Errorf("unable to start server: %w", err)
	}
	return nil
}
