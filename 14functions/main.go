package main

import "fmt"

func main() {
	// result := adder(3, 4)
	// fmt.Println("Result is: ", result)

	result := proAdder(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("Result is: ",result)
}

func adder(a int, b int) int {
	return a + b
}

//pro function -> when there is no fix number of parameter
func proAdder(values ... int) int{
	total := 0

	for _ , value := range values{
		total += value
	}

	return total
}
