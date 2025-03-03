package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //The value for format has to be learnt
	
	createdDate := time.Date(2026, time.March, 3, 14, 50, 10, 10, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}


// To build executable files in macOS or linux use the following commands:
// set GOOS=linux
// set GOARCH=amd64
// go build -o 05mytime-linux-amd64     