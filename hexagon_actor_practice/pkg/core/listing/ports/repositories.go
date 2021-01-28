package ports

import "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"

// BookRepo provides access to book repository.
type BookRepo interface {
	// AddBook saves a given book to the repository.
	GetBook(int) (domain.Book, error)
	// GetAllBeers returns all beers saved in storage.
	GetAllBooks() []domain.Book
}

// ReviewRepo provides access to review repository.
type ReviewRepo interface {
	// AddReview saves a given review.
	GetAllReviews(int) []domain.Review
}
