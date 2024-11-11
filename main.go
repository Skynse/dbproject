package main

import (
	"dbproject/db_core"
	"dbproject/web"
	"log"
)

func main() {

	// Initialize database service with proper error handling
	service, err := db_core.NewDBService()
	if err != nil {
		log.Fatalf("Failed to create database service: %v", err)
	}
	defer service.Close()

	// Create and setup server
	server := web.NewServer(service)
	server.SetupRoutes()

	// Start the server
	log.Println("Starting server on :8000")
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
