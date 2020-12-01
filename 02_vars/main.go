package main

import "fmt"

func main() {
	name, email := "Emina", "@hotmail.com"

	size := 1.3
	var age = 37
	var isCool = true
	isCool = false
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T \n", size)
}
