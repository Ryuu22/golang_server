package main

import (
	"log"

	"golang_server.dankbueno.com/internal/config"
	"golang_server.dankbueno.com/internal/repositories"
	"golang_server.dankbueno.com/internal/server"
)

func main() {
	err := config.ValidateConfig()

	if err != nil {
		log.Fatal(err)
	}

	srv := server.New() // Create new server instance

	repositories.Connect() // Connect to database

	log.Println("Starting server on port ", config.Port)
	srv.Run()
}
