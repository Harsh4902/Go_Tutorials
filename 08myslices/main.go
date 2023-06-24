package main

import (
	"fmt"
)

func main() {

	fmt.Println("-------Welcome to Intro of Slices-------")

	// var fruiteList = []string{}
	// fmt.Printf("The type of fruitlist is %T\n", fruiteList)

	// highScore := make([]int,4)

	// highScore[0] = 92
	// highScore[1] = 82
	// highScore[2] = 79
	// highScore[3] = 84
	// highScore = append(highScore, 73)
	// fmt.Println("High scores are : ", highScore)
	// sort.Ints(highScore)

	// fmt.Println("Sorted Slice is: ",highScore)
	// fmt.Println(highScore[3])


	//How to remove a value from slice based on index
	var courses = []string{"java","Spring Framework", "Spring boot", "Golang" ,"Servlets & JSP"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)
}