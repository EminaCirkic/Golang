package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/handlers/http/graphqlAPI/graph"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/handlers/http/graphqlAPI/graph/generated"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/repositories/db/mongoDb"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/repositories/files/json"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/repositories/memory"
	adding "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding"
	addport "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding/ports"
	listing "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/listing"
	listport "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/listing/ports"
)

// Type defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory

	DBMONGO Type = iota
)

func main() {

	// set up storage
	storageType := DBMONGO // this could be a flag; hardcoded here for simplicity

	var adder addport.BookService
	var lister listport.BookService
	var reviewer addport.ReviewService
	var listerr listport.ReviewService
	switch storageType {
	case Memory:
		s := new(memory.Storage)
		adder = adding.NewBookService(s) // adding "actor"
		reviewer = adding.NewReviewService(s)
		lister = listing.NewBookService(s)
		listerr = listing.NewReviewService(s)

	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()
		adder = adding.NewBookService(s) // adding "actor"
		reviewer = adding.NewReviewService(s)
		lister = listing.NewBookService(s)
		listerr = listing.NewReviewService(s)

	case DBMONGO:
		s := mongoDb.Connect()
		adder = adding.NewBookService(s) // adding "actor"
		reviewer = adding.NewReviewService(s)
		lister = listing.NewBookService(s)
		listerr = listing.NewReviewService(s)
	}

	// REST API
	// router := rest.Handler(adder, reviewer, lister, listerr)

	// fmt.Println("The book server is on now: http://localhost:3000")
	// log.Fatal(http.ListenAndServe(":3000", router))

	//GraphQL API
	graph.Handler(adder, reviewer, lister, listerr)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Print("Starting to listen 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
