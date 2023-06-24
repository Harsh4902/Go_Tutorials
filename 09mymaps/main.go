package main

import "fmt"

func main() {
	fmt.Println("Welcome to Intro of Map")

	languages := make(map[string]string)

	languages["js"] = "JavaScript"
	languages["cpp"] = "C++"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"

	fmt.Println("List of all Languages: ", languages)
	fmt.Println("js shorts for: ", languages["js"])

	delete(languages, "py")
	fmt.Println("List of all Languages: ", languages)


	//loops are intresting in Golang

	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}
}
