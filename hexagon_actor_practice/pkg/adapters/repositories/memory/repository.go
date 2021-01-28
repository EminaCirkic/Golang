package memory

import (
	"fmt"
	"strconv"
	"time"

	adding "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding/ports"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
	listing "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/listing/ports"
)

// Storage memory keeps data in memory
type Storage struct {
	books   []domain.Book
	reviews []domain.Review
}

// AddBook saves the given book to the repository
func (m *Storage) AddBook(b domain.NewBook) error {
	for _, e := range m.books {
		if b.Title == e.Title &&
			b.Author == e.Author {
			return adding.ErrDuplicate
		}
	}

	newB := domain.Book{
		ID:     strconv.Itoa(len(m.books) + 1),
		Title:  b.Title,
		Author: b.Author,
	}
	m.books = append(m.books, newB)

	return nil
}

// AddReview saves the given review in the repository
func (m *Storage) AddReview(r domain.NewReview) error {
	found := false
	for b := range m.books {
		if m.books[b].ID == r.BookID {
			found = true
		}
	}

	if found {
		created := time.Now()
		id := fmt.Sprintf("%d_%s_%s_%d", r.BookID, r.FirstName, r.LastName, created.Unix())

		newR := domain.Review{
			ID:        id,
			Created:   created,
			BookID:    r.BookID,
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Score:     r.Score,
			Text:      r.Text,
		}

		m.reviews = append(m.reviews, newR)
	} else {
		return listing.ErrNotFound
	}

	return nil
}

// GetBook returns a book with the specified ID
func (m *Storage) GetBook(id int) (domain.Book, error) {
	var book domain.Book
	bookid := strconv.Itoa(id)

	for i := range m.books {

		if m.books[i].ID == bookid {
			book.ID = m.books[i].ID
			book.Title = m.books[i].Title
			book.Author = m.books[i].Author

			return book, nil
		}
	}

	return book, listing.ErrNotFound
}

// GetAllBooks return all books
func (m *Storage) GetAllBooks() []domain.Book {
	var books []domain.Book

	for i := range m.books {

		book := domain.Book{
			ID:     m.books[i].ID,
			Title:  m.books[i].Title,
			Author: m.books[i].Author,
		}

		books = append(books, book)
	}

	return books
}

// GetAllReviews returns all reviews for a given book
func (m *Storage) GetAllReviews(bookID int) []domain.Review {
	var list []domain.Review
	bookid := strconv.Itoa(bookID)
	for i := range m.reviews {
		if m.reviews[i].BookID == bookid {
			r := domain.Review{
				ID:        m.reviews[i].ID,
				BookID:    m.reviews[i].BookID,
				FirstName: m.reviews[i].FirstName,
				LastName:  m.reviews[i].LastName,
				Score:     m.reviews[i].Score,
				Text:      m.reviews[i].Text,
				Created:   m.reviews[i].Created,
			}

			list = append(list, r)
		}
	}

	return list
}
