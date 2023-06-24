package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	//no Intheritance in golang; No super or parent

	harshvardhan := User{"Harshvardhan", "harsh@gmail.com", true, 21}
	fmt.Println(harshvardhan)
	fmt.Printf("harshvardhan details are: %+v",harshvardhan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
