package json

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"strconv"
	"time"

	adding "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding/ports"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
	listing "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/listing/ports"
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionBook identifier for the JSON collection of books
	CollectionBook = "books"
	// CollectionReview identifier for the JSON collection of reviews
	CollectionReview = "reviews"
)

// Storage stores beer data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddBook saves the given book to the repository
func (s *Storage) AddBook(b domain.NewBook) error {

	existingBeers := s.GetAllBooks()
	for _, e := range existingBeers {
		if b.Title == e.Title &&
			b.Author == e.Author {
			return adding.ErrDuplicate
		}
	}

	newB := domain.Book{
		ID:     strconv.Itoa(len(existingBeers) + 1),
		Title:  b.Title,
		Author: b.Author,
	}

	resource := newB.ID
	if err := s.db.Write(CollectionBook, resource, newB); err != nil {
		return err
	}
	return nil
}

// AddReview saves the given review in the repository
func (s *Storage) AddReview(r domain.NewReview) error {

	var book domain.Book
	if err := s.db.Read(CollectionBook, r.BookID, &book); err != nil {
		return listing.ErrNotFound
	}

	created := time.Now()
	newR := domain.Review{
		ID:        fmt.Sprintf("%d_%s_%s_%d", r.BookID, r.FirstName, r.LastName, created.Unix()),
		BookID:    r.BookID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
		Created:   created,
	}

	if err := s.db.Write(CollectionReview, newR.ID, r); err != nil {
		return err
	}

	return nil
}

// GetBook returns a book with the specified ID
func (s *Storage) GetBook(id int) (domain.Book, error) {
	var b domain.Book
	var book domain.Book

	var resource = strconv.Itoa(id)

	if err := s.db.Read(CollectionBook, resource, &b); err != nil {
		// err handling omitted for simplicity
		return book, listing.ErrNotFound
	}

	book.ID = b.ID
	book.Title = b.Title
	book.Author = b.Author

	return book, nil
}

// GetAllBooks returns all books
func (s *Storage) GetAllBooks() []domain.Book {
	list := []domain.Book{}

	records, err := s.db.ReadAll(CollectionBook)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var b domain.Book
		var book domain.Book

		if err := json.Unmarshal([]byte(r), &b); err != nil {
			// err handling omitted for simplicity
			return list
		}

		book.ID = b.ID
		book.Title = b.Title
		book.Author = b.Author

		list = append(list, book)
	}

	return list
}

// GetAllReviews returns all reviews for a given book
func (s *Storage) GetAllReviews(bookID int) []domain.Review {
	list := []domain.Review{}

	records, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	id := strconv.Itoa(bookID)
	for _, b := range records {
		var r domain.Review

		if err := json.Unmarshal([]byte(b), &r); err != nil {
			// err handling omitted for simplicity
			return list
		}

		if r.BookID == id {
			var review domain.Review

			review.ID = r.ID
			review.BookID = r.BookID
			review.FirstName = r.FirstName
			review.LastName = r.LastName
			review.Score = r.Score
			review.Text = r.Text
			review.Created = r.Created

			list = append(list, review)
		}
	}

	return list
}
