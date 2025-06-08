package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Starting JSON processing...")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
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

func DecodeJson() {
	jsonfromWeb := []byte(`
		{
			"coursename": "React JS Bootcamp",
			"Price": 299,
			"website": "Udemy",
			"tags": [
					"web dev",
					"js",
					"react"
			]
        }
	`)

	var mycourse course 

	checkValid := json.Valid(jsonfromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonfromWeb, &mycourse)
		fmt.Printf("%#v\n", mycourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonfromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type ois %T\n", key, value, myOnlineData)
	}

}
