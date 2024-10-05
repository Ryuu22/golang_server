package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang_server.dankbueno.com/internal/config"
)

type Server struct {
	Router *mux.Router
	Port   string
}

func New() *Server {
	r := mux.NewRouter()
	s := &Server{
		Router: r,
		Port:   config.Port,
	}
	s.routes()
	return s
}

func (s *Server) Run() {
	http.ListenAndServe(":"+s.Port, s.Router)
}

func (s *Server) routes() {
	s.UserRoutes()
}
