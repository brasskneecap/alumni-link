package firebase

import (
	"context"
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
		creds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if creds == "" {
			log.Println("GOOGLE_APPLICATION_CREDENTIALS is NOT set InitFirestore")
		}

		app, err := firebase.NewApp(projectCtx, nil, option.WithCredentialsFile(creds))
		if err != nil {
			initErr = err
			return
		}
		client, initErr = app.Firestore(projectCtx)
	})
	return client, initErr
}
