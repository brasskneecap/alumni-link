package firebase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	client     *firestore.Client
	once       sync.Once
	initErr    error
	projectCtx = context.Background()
)

func InitFirestore() (*firestore.Client, error) {
	once.Do(func() {
		// First check if GOOGLE_CREDENTIALS_JSON is set (Render style)
		if credsJSON := os.Getenv("GOOGLE_CREDENTIALS_JSON"); credsJSON != "" {
			log.Println("Using GOOGLE_CREDENTIALS_JSON from env")
			app, err := firebase.NewApp(
				projectCtx,
				nil,
				option.WithCredentialsJSON([]byte(credsJSON)),
			)
			if err != nil {
				initErr = fmt.Errorf("failed to init Firestore from JSON: %w", err)
				return
			}
			client, initErr = app.Firestore(projectCtx)
			return
		}

		// Fallback: use GOOGLE_APPLICATION_CREDENTIALS as a file path (local dev)
		if credsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); credsPath != "" {
			log.Println("Using GOOGLE_APPLICATION_CREDENTIALS file:", credsPath)
			app, err := firebase.NewApp(
				projectCtx,
				nil,
				option.WithCredentialsFile(credsPath),
			)
			if err != nil {
				initErr = fmt.Errorf("failed to init Firestore from file: %w", err)
				return
			}
			client, initErr = app.Firestore(projectCtx)
			return
		}

		// Neither variable is set
		initErr = errors.New("no Firestore credentials found")
	})
	return client, initErr
}
