package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	adding "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding/ports"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
	listing "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/listing/ports"
	httprouter "github.com/julienschmidt/httprouter"
)

func Handler(a adding.BookService, r adding.ReviewService, l listing.BookService, lr listing.ReviewService) http.Handler {
	router := httprouter.New()

	router.GET("/books", getBooks(l))
	router.GET("/books/:id", getBook(l))
	router.GET("/books/:id/reviews", getBookReviews(lr))

	router.POST("/addBook", addBook(a))
	router.POST("/books/:id/reviews", addBookReview(r))

	return router
}

// addBeer returns a handler for POST /beers requests
func addBook(s adding.BookService) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newBook domain.NewBook
		err := decoder.Decode(&newBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddBook(newBook)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New book added.")
	}
}

// addBeerReview returns a handler for POST /beers/:id/reviews requests
func addBookReview(s adding.ReviewService) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID := p.ByName("id")

		var newReview domain.NewReview
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&newReview); err != nil {
			http.Error(w, "Failed to parse review", http.StatusBadRequest)
		}

		newReview.BookID = ID

		s.AddBookReview(newReview)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New book review added.")
	}
}

// getBeers returns a handler for GET /beers requests
func getBooks(s listing.BookService) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetBooks()
		json.NewEncoder(w).Encode(list)
	}
}

// getBeer returns a handler for GET /beers/:id requests
func getBook(s listing.BookService) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid book ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		book, err := s.GetBook(ID)
		if err == listing.ErrNotFound {
			http.Error(w, "The book you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}
}

// getBeerReviews returns a handler for GET /beers/:id/reviews requests
func getBookReviews(s listing.ReviewService) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid book ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		reviews := s.GetBookReviews(ID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reviews)
	}
}
