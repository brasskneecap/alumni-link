package routes

import (
	"AlumniLink/api/pkg/handlers"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

func SetupRouter(client *firestore.Client) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// User Routes
	userRouter := router.PathPrefix("/user").Subrouter()
	userHandler := &handlers.UserHandler{
		Client: client,
	}
	handlers.RegisterUserRoutes(userRouter, userHandler)

	// Assignments Routes
	assignmentsRouter := router.PathPrefix("/assignments").Subrouter()
	assignmentsHandler := &handlers.AssignmentsHandler{
		Client: client,
	}
	handlers.RegisterAssignmentsRoutes(assignmentsRouter, assignmentsHandler)

	// Blasts Routes
	blastsRouter := router.PathPrefix("/blasts").Subrouter()
	blastsHandler := &handlers.BlastsHandler{
		Client: client,
	}
	handlers.RegisterBlastsRoutes(blastsRouter, blastsHandler)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})
	// Health Check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})

	return router
}
