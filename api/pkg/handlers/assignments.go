package handlers

import (
	"AlumniLink/api/pkg/stores"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

type AssignmentsHandler struct {
	Client *firestore.Client
}

func RegisterAssignmentsRoutes(router *mux.Router, h *AssignmentsHandler) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("USER ENDPOINT"))
	})
	router.HandleFunc("/{groupId}/{id}", h.GetStudentAssignments).Methods("POST", "OPTIONS")
}

func (h *AssignmentsHandler) GetStudentAssignments(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	vars := mux.Vars(r)
	groupId := vars["groupId"]
	id := vars["id"]

	fmt.Printf("Student id  = %s", id)
	// fmt.Println("Group ID: %s, Assignment ID: %s", groupId, id)
	assignments, err := stores.GetStudentAssignments(h.Client, groupId, id)

	if err != nil {
		http.Error(w, "Troule fetching assignments", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignments)
}
