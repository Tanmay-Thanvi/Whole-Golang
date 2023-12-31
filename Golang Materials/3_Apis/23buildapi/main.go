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
	CourseId    string  `json:"courseid"` // chosen for string conversion to int
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"` // hidden 
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Middleware or helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Api in Golang")

	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2",CourseName: "ReactJs",CoursePrice: 299, Author: &Author{Fullname:  "Tanmay", Website:  "google.com"}})
	courses = append(courses, Course{CourseId: "3",CourseName: "GoLang",CoursePrice: 199, Author: &Author{Fullname:  "Jay", Website:  "lco.com"}})

	
	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses/", getAllcourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	
	// listen to a part
	log.Fatal(http.ListenAndServe(":4000",r))
}

// Who is handle the situation - Controller - File

// serve Home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h3>Welcome to Api in golang</h3>"))
}

func getAllcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Println("Params : ", params)

	// loop through courses , find matching id and return response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course Found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// What about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return 
	}
	// TODO : Check only if title if duplicated 


	// generate unique id, string conversion
	// append new course into courses
	rand.Seed(time.Now().UnixNano()) // to create random no. 
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return 
}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request 
	params:= mux.Vars(r)

	// loop throgh value, get element from id , remove that element , add my ID 
	for index, course := range courses {
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
	//TODO: 
	json.NewEncoder(w).Encode("Id not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index,course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted successfully!")
			return 
		}
	}

}