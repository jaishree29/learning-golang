package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CourseId    string `json:"course_id"`
	CourseName  string `json:"course_name"`
	CoursePrice int    `json:"course_price"`
	Author      Author `json:"author"`
}

type Author struct {
	AuthorId      string `json:"author_id"`
	AuthorName    string `json:"author_name"`
	AuthorWebsite string `json:"author_website"`
}

var course []Course

func (c *Course) IsEmpty() bool {
	if c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author.AuthorId == "" && c.Author.AuthorName == "" && c.Author.AuthorWebsite == "" {
		return true
	}
	return false
}

//controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Home Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func main() {

}