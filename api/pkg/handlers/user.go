package handlers

import (
	"AlumniLink/api/pkg/stores"
	"encoding/json"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
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
	Password: "$2a$10$VdoWMA4EnEqZFM3wOrGZR.3uMeLBY3fGR2imfz8tn2n6s6OkbZhjK", // Password is "password123"
}

type UserHandler struct {
	Client *firestore.Client
}

func RegisterUserRoutes(router *mux.Router, h *UserHandler) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("USER ENDPOINT"))
	})
	router.HandleFunc("/login", h.GetUserWithCredentials).Methods("POST", "OPTIONS")
}

func (h *UserHandler) GetUserWithCredentials(w http.ResponseWriter, r *http.Request) {
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

	// Check if the user exists (mock check for now)
	user, err := stores.GetUserByEmail(h.Client, reqBody.Email)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
