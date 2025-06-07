package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://api.restful-api.dev:3000/objects?id=123&name=example"

func main() {
	fmt.Print("Starting the URL processing...\n")
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are %T\n", qparams)

	fmt.Println(qparams["id"])

	for _, value := range qparams {
		fmt.Println(value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "api.restful-api.dev",
		Path:   "/objects",
		RawQuery: "id=123&name=example",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}