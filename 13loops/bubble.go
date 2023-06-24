package main

import "fmt"

func main() {

	arr := [5]int{5, 4, 3, 2, 1}

	fmt.Println("Unsorted array: ", arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	fmt.Println("Sorted array: ", arr)
}
