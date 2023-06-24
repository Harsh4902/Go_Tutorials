package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	//no Intheritance in golang; No super or parent

	harshvardhan := User{"Harshvardhan", "harsh@gmail.com", true, 21}
	fmt.Println("Is user active: ",harshvardhan.GetStatus())
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() bool{
	return u.Status
}