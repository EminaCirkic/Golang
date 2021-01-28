package domain

//BookPayload for multiple books
type BookPayload []NewBook

//Book stuct
type Book struct {
	ID     string `json:"id" bson:"_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type NewBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
