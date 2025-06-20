package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting with modules in Go!")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	greeter()

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter(){
	fmt.Println("Hey there, I am a greeter function!")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to learning modules in Go!</h1>"))
}