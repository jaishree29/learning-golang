package main

import "fmt"

func main() {
	defer fmt.Println("world")  // This will execute after the main function completes
	defer fmt.Println("One")  // This will execute before the previous defer
	defer fmt.Println("Two") // Uses Last in first out order
	fmt.Println("hello")
	myDefer() 
}

func myDefer() {
	for i := range 5 {
		defer fmt.Println(i) 
	}
}