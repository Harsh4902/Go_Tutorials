package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("if & else in Golang")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Login count: ")

	input, _ := reader.ReadString('\n')

	loginCount, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		var result string
		if loginCount < 10 {
			result = "Regular user"
		} else if loginCount > 10 {
			result = "Watch out"
		} else {
			result = "Exactly 10 login count"
		}
		fmt.Println(result)
	}


	//special syntax, Making variable at if statement

	if num := 5; num%2 == 0{
		fmt.Println("Even")
	}else {
		fmt.Println("Odd")
	}
}
