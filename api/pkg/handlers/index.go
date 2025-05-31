package handlers

import (
	"cloud.google.com/go/firestore"
)

type FirestoreHandler struct {
	Client *firestore.Client
}
