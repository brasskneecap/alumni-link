package stores

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type User struct {
	ID       string `firestore:"-"` // we'll assign this from DocRef.ID
	Email    string `firestore:"email"`
	Password string `firestore:"password"`
	School   string `firestore:"school"`
	Name     string `firestore:"name"`
	// Add more fields as needed
}

func GetUserByEmail(client *firestore.Client, email string) (*User, error) {
	ctx := context.Background()

	iter := client.Collection("users").Where("email", "==", email).Limit(1).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		fmt.Println("No Doc Found", err)
		return nil, err // no doc found or Firestore error
	}

	var user User
	err = doc.DataTo(&user)
	if err != nil {
		return nil, err
	}

	user.ID = doc.Ref.ID // set the document ID if you want it
	fmt.Println("User Found", user)
	return &user, nil
}
