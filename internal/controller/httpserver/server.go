package httpserver

import (
	"net/http"
)

type server struct {
	s http.Server
}

func NewServer(addr string, handler http.Handler) *server {
	return &server{
		s: http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *server) Start() error {
	return s.s.ListenAndServe()
}

func (s *server) Stop() error {
	return s.s.Shutdown(nil)
}
