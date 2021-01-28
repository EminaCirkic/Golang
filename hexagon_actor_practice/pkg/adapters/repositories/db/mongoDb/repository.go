package mongoDb

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding/ports"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Storage struct {
	client *mongo.Client
}

func Connect() *Storage {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{
		client: client,
	}
}

// AddBook saves the given book to the repository
func (s *Storage) AddBook(b domain.NewBook) error {

	existingBeers := s.GetAllBooks()
	for _, e := range existingBeers {
		if b.Title == e.Title &&
			b.Author == e.Author {
			return ports.ErrDuplicate
		}
	}

	newB := domain.NewBook{
		Title:  b.Title,
		Author: b.Author,
	}
	collection := s.client.Database("books").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, newB)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

// AddReview saves the given review in the repository
func (s *Storage) AddReview(r domain.NewReview) error {

	created := time.Now()
	newR := domain.NewReview{
		BookID:    r.BookID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
		Created:   created,
	}

	collection := s.client.Database("books").Collection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, newR)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
}

// GetBook returns a book with the specified ID
func (s *Storage) GetBook(id int) (domain.Book, error) {

	var newID = strconv.Itoa(id)
	ObjectID, err := primitive.ObjectIDFromHex(newID)
	if err != nil {
		log.Fatal(err)
	}
	collection := s.client.Database("books").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})

	book := domain.Book{}
	res.Decode(&book)

	return book, nil
}

// GetAllBooks returns all books
func (s *Storage) GetAllBooks() []domain.Book {
	list := []domain.Book{}
	collection := s.client.Database("books").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	records, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for records.Next(ctx) {
		var book domain.Book

		err := records.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, book)
	}

	return list
}

// GetAllReviews returns all reviews for a given book
func (s *Storage) GetAllReviews(bookID int) []domain.Review {
	list := []domain.Review{}

	collection := s.client.Database("books").Collection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	records, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for records.Next(ctx) {
		var review domain.Review

		err := records.Decode(&review)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, review)
	}

	return list
}
