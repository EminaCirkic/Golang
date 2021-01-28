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

	// Failed means processing did not finish successfully
	Failed

	// We could also have a Queued Event which would mean queued for processing
	Queued
)

func (e Event) GetMeaning() string {
	if e == Done {
		return "Done"
	}

	if e == Failed {
		return "Failed"
	}

	if e == Queued {
		return "Processing.."
	}

	return "Unknown result"
}

// ErrNotFound is used when a book could not be found.
var ErrNotFound = errors.New("book not found")

type BookService interface {
	GetBook(int) (domain.Book, error)
	GetBooks() []domain.Book
}

type ReviewService interface {
	GetBookReviews(int) []domain.Review
}
