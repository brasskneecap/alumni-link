package handlers

import (
	"AlumniLink/api/pkg/stores"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

type BlastsHandler struct {
	Client *firestore.Client
}

func RegisterBlastsRoutes(router *mux.Router, h *BlastsHandler) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("USER ENDPOINT"))
	})
	router.HandleFunc("/{groupId}", h.GetBlasts).Methods("GET", "POST", "OPTIONS")
}

func (h *BlastsHandler) GetBlasts(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	vars := mux.Vars(r)
	groupId := vars["groupId"]

	fmt.Printf("Group id  = %s", groupId)
	// fmt.Println("Group ID: %s, Assignment ID: %s", groupId, id)
	blasts, err := stores.GetBlasts(h.Client, groupId)

	if err != nil {
		http.Error(w, "Troule fetching blasts", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blasts)
}
