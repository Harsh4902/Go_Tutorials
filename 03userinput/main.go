package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to UserInput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	//comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Your Name is:", input)

}
