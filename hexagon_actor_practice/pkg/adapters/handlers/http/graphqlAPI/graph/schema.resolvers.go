package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"

	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/handlers/http/graphqlAPI/graph/generated"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/handlers/http/graphqlAPI/graph/model"
	adding "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding/ports"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
	listing "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/listing/ports"
)

var listr []domain.Book
var listrR []domain.Review
var listi listing.BookService

func Handler(a adding.BookService, r adding.ReviewService, l listing.BookService, lr listing.ReviewService) {
	listr = l.GetBooks()
	listi = l
}

func (r *queryResolver) Getbook(ctx context.Context, id *string) (*model.Book, error) {
	ID, err := strconv.Atoi(*id)
	Id := *id
	if listr[ID-1].ID == Id {
		book := listr[ID-1]
		return &model.Book{
			ID:     id,
			Title:  &book.Title,
			Author: &book.Author,
		}, nil
	} else {
		log.Fatal(err, "Book not found")
	}
	return nil, nil
}

func (r *queryResolver) GetBooks(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	list := listi.GetBooks()
	for i := range list {
		b := list[i]
		books = append(books, &model.Book{
			ID:     &b.ID,
			Title:  &b.Title,
			Author: &b.Author,
		})
	}
	return books, nil
	//
	// for i := range listr {
	// 	var book *model.Book
	// 	b := listr[i]
	// 	books = append(books, &model.Book{
	// 		ID:     &b.ID,
	// 		Title:  &b.Title,
	// 		Author: &b.Author,
	// 	})
	// }

	// return books, nil
	//panic(fmt.Errorf("not implemented"))
	return nil, nil

}

func (r *queryResolver) GetBookReview(ctx context.Context, id *string) ([]*model.Review, error) {
	return nil, nil

}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
