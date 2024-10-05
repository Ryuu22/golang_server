package server

import (
	"golang_server.dankbueno.com/internal/handlers"
)

func (s *Server) UserRoutes() {
	s.Router.HandleFunc("/api/v1/user", handlers.CreateUser).Methods("POST")
	s.Router.HandleFunc("/api/v1/login", handlers.LogIn).Methods("POST")
	s.Router.HandleFunc("/api/v1/user", handlers.GetUser).Methods("GET")
}
