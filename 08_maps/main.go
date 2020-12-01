package main

import "fmt"

func main() {
	//Long way
	//Define map
	//emails := make(map[string]string)

	//Assign key
	//emails["Bob"] = "@gmail.com"
	//emails["Sharon"] = "@hotmail.com"
	//emails["Mike"] = "@mail.com"

	//Short way
	emails := map[string]string{"Bob": "@gmail.com", "Sharon": "@Hotmail.com"}

	delete(emails, "Bob")
	fmt.Println(emails)

}
