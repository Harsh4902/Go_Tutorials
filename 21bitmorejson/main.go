package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {

	fmt.Println("Welcome to Json")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson(){

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev","js"}},
		{"Spring Bootcamp", 100, "LearnCodeOnline.in", "xyz567", []string{"spring","java"}},
		{"Android Bootcamp", 499, "LearnCodeOnline.in", "3456", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)

} 

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"Name": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Password": "abc123",
		"Tags": [
				"web-dev",
				"js"
		]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	//some 

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	//fmt.Print("%#v\n", myOnlineData)

	for index,value := range myOnlineData{
		fmt.Printf("%v : %v\n",index,value)
	}
}
