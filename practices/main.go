package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myUrl string = "https://api.restful-api.dev/objects"

func main() {
	makeGetRequest()
	makePostRequest()
}

func makeGetRequest() {
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
	defer response.Body.Close()
}

func makePostRequest() {
	requestBody := strings.NewReader(`
		{
			"name": "Apple MacBook Pro 16",
			"data": {
			"year": 2019,
			"price": 1849.99,
			"CPU model": "Intel Core i9",
			"Hard disk size": "1 TB"
			}
	   }
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
	defer response.Body.Close()
}


