package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza")

	//comma ok || error err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is: %T\n", input)

	//using scan
	fmt.Println("Enter your name:")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Thank you %s, for rating our pizza!", name)

	//The reason for NOT using SCAN method is that it only reads upto a white space character so if you enter a name say, Jaishree Tiwari, then it will give an error

	
}