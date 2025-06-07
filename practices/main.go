package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myUrl string = "https://api.restful-api.dev/objects"

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

func main() {
	makeGetRequest()
}