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

type UserHandler struct {
	Client *firestore.Client
}

func RegisterUserRoutes(router *mux.Router, h *UserHandler) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("USER ENDPOINT"))
	})
	router.HandleFunc("/login", h.GetUserWithCredentials).Methods("POST", "OPTIONS")
	router.HandleFunc("/users", h.GetUsers).Methods("POST", "OPTIONS")
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
		http.Error(w, "Internal Server Error", http.StatusNotFound)
		return
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var reqBody struct {
		Groups []string `json:"groups"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the user exists (mock check for now)
	users, err := stores.GetUsers(h.Client, reqBody.Groups)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
