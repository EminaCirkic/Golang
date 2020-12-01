package main

import "fmt"

func main() {
	//Arrays
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declare and assign
	//fruitArr := [2]string{"Apple", "Orange"}

	fruitArrSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitArrSlice))
	fmt.Println(fruitArrSlice[1:2])

}
