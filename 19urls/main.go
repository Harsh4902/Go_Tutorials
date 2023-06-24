package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&pamentid=ghjb4567g"

func main() {

	fmt.Println("Welcome to handling URL in Golang")
	fmt.Println(myUrl)

	//parsing url
	result, _ := url.Parse(myUrl)

	qparams := result.Query()

	for _,value := range qparams{
		fmt.Println(value)
	}
}
