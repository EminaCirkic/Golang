package main

import (
	"fmt"
	"time"

	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/adapters/repositories/files/json"
	adding "github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/adding"
	"github.com/eminacirkic/go_crash_course/hexagon_actor_practice/pkg/core/domain"
)

type Message interface{}

func main() {

	// error handling omitted for simplicity
	s, _ := json.NewStorage()

	// create the available services
	adder := adding.NewBookService(s)      // adding "actor"
	reviewer := adding.NewReviewService(s) // reviewing "actor"

	resultsBook := adder.AddSampleBooks(domain.DefaultBooks)
	resultsReview := reviewer.AddSampleReviews(domain.DefaultReviews)

	go func() {
		for result := range resultsBook {
			fmt.Printf("Added sample book with result %s.\n", result.GetMeaning()) // human-friendly
		}
	}()

	go func() {
		for result := range resultsReview {
			fmt.Printf("Added sample review with result %s.\n", result.GetMeaning()) // machine-friendly
		}
	}()

	// main could have its own "mailbox" exposed, for example an HTTP endpoint,
	// so we could be waiting here for more sample data to be added
	// (but we'll just exit for simplicity)

	time.Sleep(2 * time.Second) // this is here just to get the output from goroutines printed

	fmt.Println("No more data to add!")
}
