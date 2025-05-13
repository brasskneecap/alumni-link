package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `json:"name"`
	School   string `json:"school"`
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Mock user for testing
var mockUser = User{
	Name:     "Test User",
	School:   "Test School",
	ID:       "123",
	Email:    "test@example.com",
	Password: "$2a$10$V88UKhSytAlx.Us2zFzc7uulzIo2ICK54nW1TgA4m6zm4zw3QIbji", // Password is "password123"
}

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login", GetUserWithCredentials).Methods("POST", "OPTIONS")
}

func GetUserWithCredentials(w http.ResponseWriter, r *http.Request) {
	// Parse the request body

	var reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Trim the email to avoid simple user errors
	reqBody.Email = strings.TrimSpace(reqBody.Email)

	fmt.Println("reqBody ", reqBody.Email)
	fmt.Println("mockUser.Email ", mockUser.Email)
	// Check if the user exists (mock check for now)
	if reqBody.Email != mockUser.Email {
		fmt.Println("Email doesnt match")
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(mockUser.Password), []byte(reqBody.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Respond with the user details (excluding password)
	user := struct {
		Name   string `json:"name"`
		School string `json:"school"`
		ID     string `json:"id"`
		Email  string `json:"email"`
	}{
		Name:   mockUser.Name,
		School: mockUser.School,
		ID:     mockUser.ID,
		Email:  mockUser.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
