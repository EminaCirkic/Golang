package adding

import (
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding/ports"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
)

type bookservice struct {
	bookRepository ports.BookRepo
}

type reviewservice struct {
	reviewRepository ports.ReviewRepo
}

// NewBookService creates an adding service with the necessary dependencies
func NewBookService(bookRepository ports.BookRepo) *bookservice {
	return &bookservice{
		bookRepository: bookRepository,
	}
}

// NewReviewService creates an adding service with the necessary dependencies
func NewReviewService(reviewRepository ports.ReviewRepo) *reviewservice {
	return &reviewservice{
		reviewRepository: reviewRepository,
	}
}

// AddBook adds the given book(s) to the database
func (s *bookservice) AddBook(b ...domain.NewBook) {

	// any validation can be done here

	for _, book := range b {
		_ = s.bookRepository.AddBook(book) // error handling omitted for simplicity
	}
}

// AddSampleBooks adds some sample books to the database
func (s *bookservice) AddSampleBooks(data domain.BookPayload) <-chan ports.Event {
	results := make(chan ports.Event)

	go func() {
		defer close(results)

		for _, b := range data {
			err := s.bookRepository.AddBook(b)
			if err != nil {
				if err == ports.ErrDuplicate {
					// forgive the naughty error type checking above...
					results <- ports.BookAlreadyExists
					continue
				}
				results <- ports.Failed
				continue
			}

			results <- ports.Done
		}
	}()

	return results
}

// AddBookReview saves a new book review in the database
func (s *reviewservice) AddBookReview(r domain.NewReview) {
	_ = s.reviewRepository.AddReview(r) // error handling omitted for simplicity
}

// AddSampleReviews adds some sample reviews to the database
func (s *reviewservice) AddSampleReviews(data domain.ReviewPayload) <-chan ports.Event {
	results := make(chan ports.Event)

	go func() {
		defer close(results)

		for _, r := range data {
			err := s.reviewRepository.AddReview(r)
			if err != nil {
				results <- ports.Failed
			}
			results <- ports.Done
		}
	}()

	return results
}
