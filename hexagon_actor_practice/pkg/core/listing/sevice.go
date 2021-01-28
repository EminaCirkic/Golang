package listing

import (
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/listing/ports"
)

type bookservice struct {
	bookRepository ports.BookRepo
}

type reviewservice struct {
	reviewRepository ports.ReviewRepo
}

// NewBookService creates an listing service with the necessary dependencies
func NewBookService(bookRepository ports.BookRepo) *bookservice {
	return &bookservice{
		bookRepository: bookRepository,
	}
}

// NewReviewService creates an listing service with the necessary dependencies
func NewReviewService(reviewRepository ports.ReviewRepo) *reviewservice {
	return &reviewservice{
		reviewRepository: reviewRepository,
	}
}

// GetBook returns a book
func (s *bookservice) GetBook(id int) (domain.Book, error) {
	return s.bookRepository.GetBook(id)
}

// GetBooks returns all books
func (s *bookservice) GetBooks() []domain.Book {
	return s.bookRepository.GetAllBooks()
}

// GetBookReviews returns all requests for a book
func (s *reviewservice) GetBookReviews(bookID int) []domain.Review {
	return s.reviewRepository.GetAllReviews(bookID)
}
