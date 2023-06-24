package main

import "fmt"

func main() {

	defer fmt.Println("Parmar")
	defer fmt.Println("Harshvardhan")
	defer fmt.Println("Hello world")
	fmt.Println("Welcome to intro of defer in Golang")
	myDefer()
}

func myDefer(){

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}