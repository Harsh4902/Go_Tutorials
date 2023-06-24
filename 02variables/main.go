package main

import "fmt"

func main() {
	var username string = "Harshvardhan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable is of type: %T \n", isTrue)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.455434566547
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	//default values and some aliases
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit types
	var hello = "Hello World !"
	fmt.Println(hello)
	fmt.Printf("Variable is of type: %T \n", hello)

	//no var style
	name := "Harshvardhan"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)
	
}
