package app

import (
	"log"
	"net/http"
)

// Server is the top level data structure.
type Server struct {
	*DB
}

// NewServer is a factory method for creating a new server.
func NewServer() *Server {
	return &Server{
		NewDB(),
	}
}

// Start the API server application
func (s *Server) Start() {
	r := s.InitRouter()
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
