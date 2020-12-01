package main

import "fmt"

//Reason to use a pointer is that you might have to pass alot of data at an address
// So to increase perfomance pass the address of the data instead of the data in order
// to increase performance
func main() {
	a := 5
	//pointer of a
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)

	//Change cal with pointer
	*b = 10
	fmt.Println(a)

}
