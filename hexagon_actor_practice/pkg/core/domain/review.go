package domain

import (
	"time"
)

//ReviewPayload for multiple reviews
type ReviewPayload []NewReview

//Review of a book with ID
type Review struct {
	ID        string    `json:"id" bson:"_id"`
	BookID    string    `json:"book_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}

type NewReview struct {
	BookID    string    `json:"book_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}
