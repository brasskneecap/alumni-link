package main

import (
	"AlumniLink/api/pkg/firebase"
	"AlumniLink/api/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	credsJSON := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credsJSON == "" {
		// Set credential path if not using environment variables already
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "config/testcreds.json")
	}

	// Initialize Firestore client
	fsClient, err := firebase.InitFirestore()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}
	defer fsClient.Close()

	router := routes.SetupRouter(fsClient)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true, // Turn this on for detailed logs
	})

	handler := c.Handler(router)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
