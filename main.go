package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	InitDB()
	defer CloseDB()

	// Setup routes
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/users", UsersHandler)

	// Start server
	port := ":8080"
	fmt.Println("🚀 Server starting on http://localhost" + port)
	fmt.Println("📍 Endpoints:")
	fmt.Println("   GET http://localhost:8080/health")
	fmt.Println("   GET http://localhost:8080/users")

	log.Fatal(http.ListenAndServe(port, nil))
}
