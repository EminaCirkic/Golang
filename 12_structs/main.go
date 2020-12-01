package main

import (
	"fmt"
	"strconv"
)

//Define person struct, instead of classes, methods associated with stuct are
// pointer and value receivers a.k.a getters and setters
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

//Value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//Ponter receiver, changes value
func (p *Person) hasBirthday() {
	p.age++
}

//Pointer receiver, changes value
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}

}
func main() {

	//Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	//Get single field
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person2 := Person{"Bob", "Johnson", "New York", "m", 30}
	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person1.greet(), person2.greet())

}
