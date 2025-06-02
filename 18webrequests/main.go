package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.restful-api.dev/objects"

func main() {
	fmt.Println("Starting web requests...")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println(("Response content:\n" + content))

	defer response.Body.Close()
	fmt.Printf("Response status: %s\n", response.Status)
}