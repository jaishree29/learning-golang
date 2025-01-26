package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int 
	// fmt.Println(ptr)
	fmt.Println("Enter an integer:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')


	var ptr = &input //& is used for reference
	fmt.Println("The value of pointer for the entered number is,", ptr)
	fmt.Println("The value of pointer is,", *ptr)


	myNumber := 23
	var newPtr = &myNumber
	*newPtr = *newPtr * 2 // Here we are multiplying the value of newPtr
	fmt.Println(myNumber)
}