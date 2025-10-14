package stores

import (
	"AlumniLink/api/pkg/models"
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type SubmissionContent struct {
	FilePath *string `firestore:"file_path,omitempty" json:"filePath,omitempty"`
	Text     *string `firestore:"text,omitempty" json:"text,omitempty"`
	URL      *string `firestore:"url,omitempty" json:"url,omitempty"`
}

type Submission struct {
	AssignmentID string            `firestore:"assignment_id" json:"assignmentId"`
	StudentID    string            `firestore:"student_id" json:"studentId"`
	SubmittedAt  time.Time         `firestore:"submitted_at" json:"submittedAt"`
	Status       string            `firestore:"status" json:"status"`
	Type         string            `firestore:"type" json:"type"`
	Feedback     string            `firestore:"feedback" json:"feedback"`
	Content      SubmissionContent `firestore:"content" json:"content"`
}

type Assignment struct {
	ID             string    `firestore:"-" json:"id"` // firestore:"-" means not stored in Firestore
	Name           string    `firestore:"name" json:"name"`
	Description    string    `firestore:"description" json:"description"`
	GroupID        string    `firestore:"group_id" json:"groupId"`
	MentorID       string    `firestore:"mentor_id" json:"mentorId"`
	AllowedContent []string  `firestore:"allowed_content" json:"allowedContent"`
	CreatedAt      time.Time `firestore:"created_at" json:"createdAt"`
	DueDate        time.Time `firestore:"due_date" json:"dueDate"`
	PublishDate    time.Time `firestore:"publish_date" json:"publishDate"`
}

type StudentAssignments struct {
	StudentID   string                     `json:"studentId"`
	Assignments []AssignmentWithSubmission `json:"assignments"`
}

// {
// "name":"test ",
// "description":"test test test",
// "groupId":"p28tMvMQBZ81cxjACsMx"
// "mentorId":"p28tMvMQBZ81cxjACsMx",
// "allowedContent":["Text"],
// "dueDate":"2025-08-02T23:47:00.000Z",
// "publishDate":"2025-07-31T06:00:00.000Z",
// }
type AssignmentWithSubmission struct {
	Assignment
	Submission *Submission `json:"submission,omitempty"`
}

func GetStudentAssignments(client *firestore.Client, groupId string, studentId string) ([]AssignmentWithSubmission, error) {
	ctx := context.Background()

	var results []AssignmentWithSubmission

	// 1. Fetch all assignments 9CrmcRDbrBW4gldZFiV1 9CrmcRDbrBW4gldZFiV1
	assignmentsIter := client.Collection("assignments").Where("group_id", "==", groupId).OrderBy("due_date", firestore.Asc).Documents(ctx)

	var assignments []Assignment
	for {
		doc, err := assignmentsIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}

		var assignment Assignment
		err = doc.DataTo(&assignment)
		if err != nil {
			return nil, err
		}
		assignment.ID = doc.Ref.ID
		assignments = append(assignments, assignment)
	}

	// 2. Fetch all submissions for the student (filter by studentID, optionally filter by assignmentIDs)
	submissionsIter := client.Collection("submissions").
		Where("student_id", "==", studentId).
		Documents(ctx)

	submissionsMap := make(map[string]*Submission) // key = assignmentID

	count := 0
	for {
		doc, err := submissionsIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, err
		}

		var sub Submission
		err = doc.DataTo(&sub)
		if err != nil {
			fmt.Println("Error with DataTo")
			return nil, err
		}

		// Map submissions by assignment ID
		submissionsMap[sub.AssignmentID] = &sub
		count++
	}
	// 3. Combine into AssignmentWithSubmission
	for _, assignment := range assignments {
		submission := submissionsMap[assignment.ID]
		results = append(results, AssignmentWithSubmission{
			Assignment: assignment,
			Submission: submission,
		})
	}

	return results, nil
}

func GetGroupAssignments(client *firestore.Client, groupId, mentorId string) ([]StudentAssignments, error) {
	ctx := context.Background()

	// 1. Get all students in the mentor’s group(s)
	usersIter := client.Collection("users").
		Where("groups", "array-contains", groupId).
		Where("mentor_id", "==", mentorId).
		Documents(ctx)

	var students []string
	for {
		doc, err := usersIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, fmt.Errorf("error fetching users: %w", err)
		}
		students = append(students, doc.Ref.ID)
	}

	if len(students) == 0 {
		return []StudentAssignments{}, nil
	}

	// 2. Fetch all assignments for the group
	assignmentsIter := client.Collection("assignments").
		Where("group_id", "==", groupId).
		OrderBy("due_date", firestore.Asc).
		Documents(ctx)

	var assignments []Assignment
	for {
		doc, err := assignmentsIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, fmt.Errorf("error fetching assignments: %w", err)
		}

		var assignment Assignment
		if err := doc.DataTo(&assignment); err != nil {
			return nil, err
		}
		assignment.ID = doc.Ref.ID
		assignments = append(assignments, assignment)
	}

	// 3. Fetch all submissions for these students (across all assignments)
	submissionsIter := client.Collection("submissions").
		Where("student_id", "in", students).
		Documents(ctx)

	submissionsByStudent := make(map[string]map[string]*Submission) // studentID -> assignmentID -> submission
	for {
		doc, err := submissionsIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, fmt.Errorf("error fetching submissions: %w", err)
		}
		var sub Submission
		if err := doc.DataTo(&sub); err != nil {
			return nil, err
		}
		if submissionsByStudent[sub.StudentID] == nil {
			submissionsByStudent[sub.StudentID] = make(map[string]*Submission)
		}
		submissionsByStudent[sub.StudentID][sub.AssignmentID] = &sub
	}

	// 4. Combine results per student
	var results []StudentAssignments
	for _, studentID := range students {
		var studentAssignments []AssignmentWithSubmission
		for _, assignment := range assignments {
			sub := submissionsByStudent[studentID][assignment.ID]
			studentAssignments = append(studentAssignments, AssignmentWithSubmission{
				Assignment: assignment,
				Submission: sub,
			})
		}
		results = append(results, StudentAssignments{
			StudentID:   studentID,
			Assignments: studentAssignments,
		})
	}

	fmt.Print("Get group assignments", results)
	return results, nil
}

func CreateAssignment(ctx context.Context, client *firestore.Client, assignment *models.Assignment) (*models.Assignment, error) {
	if assignment.CreatedAt.IsZero() {
		assignment.CreatedAt = time.Now()
	}

	docRef, _, err := client.Collection("assignments").Add(ctx, assignment)
	if err != nil {
		return nil, err
	}

	assignment.ID = docRef.ID
	return assignment, nil
}
