package routes

import (
	"AlumniLink/api/pkg/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// User Routes
	userRouter := router.PathPrefix("/user").Subrouter()
	handlers.RegisterUserRoutes(userRouter)

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
