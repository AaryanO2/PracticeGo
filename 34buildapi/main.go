package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model in a file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware,helper -seperate file
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

//controller -seperate file

// serve home
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to courses API</h1>"))
}

func getALlCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GeT one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // returns map
	fmt.Println(params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("NO COURSE FOUND FOR THIS ID")
}
func main() {
	fmt.Println("MAIN API")
	r := mux.NewRouter()
	courses = append(courses,
		Course{
			CourseId:    "1",
			CourseName:  "Introduction to Go",
			CoursePrice: 100,
			Author:      &Author{FullName: "John Doe", Website: "johndoe.com"},
		},
		Course{
			CourseId:    "2",
			CourseName:  "Advanced Go Programming",
			CoursePrice: 150,
			Author:      &Author{FullName: "Jane Smith", Website: "janesmith.com"},
		},
		Course{
			CourseId:    "3",
			CourseName:  "Web Development with Go",
			CoursePrice: 200,
			Author:      &Author{FullName: "Michael Johnson", Website: "michaeljohnson.com"},
		},
		Course{
			CourseId:    "4",
			CourseName:  "Data Structures in Go",
			CoursePrice: 120,
			Author:      &Author{FullName: "Emily Davis", Website: "emilydavis.com"},
		},
		Course{
			CourseId:    "5",
			CourseName:  "Microservices with Go",
			CoursePrice: 180,
			Author:      &Author{FullName: "Sarah Wilson", Website: "sarahwilson.com"},
		},
	)

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getALlCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	//challenge generate unique id
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("ID NOT FOUND")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode("Removed successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("ID NOT FOUND")
}
