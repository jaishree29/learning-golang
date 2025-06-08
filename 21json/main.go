package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Starting JSON processing...")

	mycourse := []course{
		{
			"React JS Bootcamp",
			299,
			"Udemy",
			"reactjs123",
			[]string{"web dev", "js", "react"},
		},
		{
			"Full Stack Web Development",
			499,
			"Coursera",
			"fullstack456",
			[]string{"web dev", "js", "react", "node"},
		},
		{
			"Data Science with Python",
			399,
			"edX",
			"datascience789",
			[]string{},
		},
	}

	finalJson, err := json.MarshalIndent(mycourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}
