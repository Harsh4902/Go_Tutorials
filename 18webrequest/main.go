package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Responce is of type: %T", responce)

	defer responce.Body.Close() //caller's responsibility to close the connection

	dataByte,err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)

	fmt.Println(content)
}
