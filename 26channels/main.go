package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in Golang")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}


	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	go func(ch chan int,wg *sync.WaitGroup){
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh,wg)

	go func(ch chan int,wg *sync.WaitGroup){
		myCh <- 5
		myCh <- 6
		wg.Done()
	}(myCh,wg)

	wg.Wait()
}
