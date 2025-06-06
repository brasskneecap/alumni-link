package stores

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

type User struct {
	ID       string   `firestore:"-" json:"id"` // Not stored in Firestore, but included in JSON
	Email    string   `firestore:"email" json:"email"`
	Password string   `firestore:"password" json:"password"`
	School   string   `firestore:"school" json:"school"`
	Name     string   `firestore:"name" json:"name"`
	Role     string   `firestore:"role" json:"role"`
	MentorId string   `firestore:"mentor_id" json:"mentorId"`
	Groups   []string `firestore:"groups" json:"groups"`
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
