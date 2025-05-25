package firebase

import (
	"context"
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
		creds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if creds == "" {
			creds = "config/test-project-252419-db6507ef65e2.json" // fallback default path
		}
		val := option.WithCredentialsFile(creds)
		app, err := firebase.NewApp(projectCtx, nil, val)
		if err != nil {
			initErr = err
			return
		}
		client, initErr = app.Firestore(projectCtx)
	})
	return client, initErr
}
