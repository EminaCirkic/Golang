package ports

import (
	"errors"

	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
)

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota

	// BookAlreadyExists means the given book is a duplicate of an existing one
	BookAlreadyExists

	// Failed means processing did not finish successfully
	Failed

	// We could also have a Queued Event which would mean queued for processing
	Queued
)

func (e Event) GetMeaning() string {
	if e == Done {
		return "Done"
	}

	if e == BookAlreadyExists {
		return "Duplicate book"
	}

	if e == Failed {
		return "Failed"
	}

	if e == Queued {
		return "Processing.."
	}

	return "Unknown result"
}

//ErrDuplicate means book already exists
var ErrDuplicate = errors.New("book already exists")

type BookService interface {
	AddBook(...domain.NewBook)
	AddSampleBooks(domain.BookPayload) <-chan Event
}

type ReviewService interface {
	AddBookReview(domain.NewReview)
	AddSampleReviews(domain.ReviewPayload) <-chan Event
}
