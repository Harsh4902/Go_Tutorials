package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to file processing in Golang")

	content := "This needs to go in file - Java"

	file,err := os.Create("./mygofile.txt")

	if err != nil{
		panic(err)
	}
	
	length,err := io.WriteString(file,content)

	if err != nil{
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mygofile.txt")

}

func readFile(fileName string){
	dataByte, err := ioutil.ReadFile(fileName)

	if err != nil{
		panic(err)
	}

	fmt.Println("Text data inside the file is:\n",string(dataByte))
}
