package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrize int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	//seeding values in DB

	courses = append(courses, Course{CourseId: "01",CourseName: "Core Java",
	CoursePrize: 199,Author: &Author{FullName: "Harshvardhan Parmar",Website: "learnjava.in",}})
	courses = append(courses, Course{CourseId: "02",CourseName: "Advanced Java",
	CoursePrize: 299,Author: &Author{FullName: "Harshvardhan Parmar",Website: "learnjava.in",}})

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	

	//listen to a port
	log.Fatal(http.ListenAndServe(":8080",r))
}

//controller - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Harshvardhan Parmar in GOlang"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get aLl courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses -> find matching id -> return the responce

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("content-type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what if: data is -> {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate a unique id
	//convert into string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Update one course")
	w.Header().Set("content-type", "application/json")

	//first grab id from req
	params := mux.Vars(r)

	//loop, id, remove, add with my ID 
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: send a response when id is not found
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses{

		if course.CourseId == params["id"]{

			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is Deleted")
			return
		}

	}

	json.NewEncoder(w).Encode("No course found with given id")
	return
}