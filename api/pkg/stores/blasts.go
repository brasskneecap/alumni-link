package stores

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Blast struct {
	ID          string    `firestore:"-" json:"id"` // Not stored in Firestore, but included in JSON
	Title       string    `firestore:"title" json:"title"`
	Description string    `firestore:"description" json:"description"`
	GroupID     string    `firestore:"group_id" json:"groupId"`
	CreatedBy   string    `firestore:"created_by" json:"createdBy"`
	CreatedAt   time.Time `firestore:"created_at" json:"createdAt"`
	// Add more fields as needed
}

func GetBlasts(client *firestore.Client, groupId string) ([]Blast, error) {
	ctx := context.Background()

	blastsIter := client.Collection("blasts").Where("group_id", "==", groupId).Where("deleted_at", "==", nil).Documents(ctx)

	var blasts []Blast
	for {
		doc, err := blastsIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}

		var blast Blast
		err = doc.DataTo(&blast)
		if err != nil {
			return nil, err
		}
		blast.ID = doc.Ref.ID
		blasts = append(blasts, blast)
	}

	return blasts, nil
}
