package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to web requests")
	//PerfromGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerfromGetRequest(){
	const myUrl = "http://localhost:8080/student"

	responce, err := http.Get(myUrl)

	if err != nil {
		fmt.Println("Err")
		panic(err)
	}

	defer responce.Body.Close()

	fmt.Println("Status code: ",responce.StatusCode)
	fmt.Println("Content length: ",responce.ContentLength)

	var responceString strings.Builder
	content, _ := ioutil.ReadAll(responce.Body)
	byteCount, _ := responceString.Write(content)
	fmt.Println(responceString.String())

	fmt.Println("Byte Count is: ",byteCount)

	//fmt.Println(string(content))

}

func PerformPostJsonRequest(){

	const myUrl = "http://localhost:8080/create"

	requestBody := strings.NewReader(`
		{
			"firstName":"Hars",
			"lastName":"Parmar",
			"id":2
		}
	`)

	responce, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
	var responceString strings.Builder
	content, _ := ioutil.ReadAll(responce.Body)
	responceString.Write(content)
	fmt.Println(responceString.String())
	

}

func PerformPostFormRequest(){
	const myUrl = "http://localhost:8080/form"

	//formdata
	data := url.Values{}

	data.Add("firstName","Harshvardhan")
	data.Add("lastName","Parmar")
	data.Add("id","1")

	responce, err := http.PostForm(myUrl,data)
	
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
	var responceString strings.Builder
	content, _ := ioutil.ReadAll(responce.Body)
	responceString.Write(content)
	fmt.Println(responceString.String())
}