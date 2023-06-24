package main

import "fmt"

func main() {
	fmt.Println("Welcome to explaination of pointer")

	//declaring pointer

	// var ptr *int
	// fmt.Println("Value of pointer is ",ptr)

	myVar := 23
	var ptr = &myVar

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)

	myVar += 5
	fmt.Println(*ptr)
}