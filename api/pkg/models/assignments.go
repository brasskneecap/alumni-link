package models

import "time"

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
	ID             string    `firestore:"-" json:"id"`
	Name           string    `firestore:"name" json:"name"`
	Description    string    `firestore:"description" json:"description"`
	GroupID        string    `firestore:"group_id" json:"groupId"`
	MentorID       string    `firestore:"mentor_id" json:"mentorId"`
	AllowedContent []string  `firestore:"allowed_content" json:"allowedContent"`
	CreatedAt      time.Time `firestore:"created_at" json:"createdAt"`
	DueDate        time.Time `firestore:"due_date" json:"dueDate"`
	PublishDate    time.Time `firestore:"publish_date" json:"publishDate"`
}

type NewAssignment struct {
	Name           string    `firestore:"name" json:"name"`
	Description    string    `firestore:"description" json:"description"`
	GroupID        string    `firestore:"group_id" json:"groupId"`
	MentorID       string    `firestore:"mentor_id" json:"mentorId"`
	AllowedContent []string  `firestore:"allowed_content" json:"allowedContent"`
	CreatedAt      time.Time `firestore:"created_at" json:"createdAt"`
	DueDate        time.Time `firestore:"due_date" json:"dueDate"`
	PublishDate    time.Time `firestore:"publish_date" json:"publishDate"`
}
