package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "API is running",
	})
}

// UsersHandler handles GET requests for all users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := GetAllUsers()
	if err != nil {
		log.Println("Error fetching users:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
