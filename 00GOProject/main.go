package main

import (
	"fmt"
	"log"
	"net/http"

	"gobackend/cmd/api"
)

func main() {
	fmt.Println("Starting Go Backend Server...")

	server := api.NewServer()

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
