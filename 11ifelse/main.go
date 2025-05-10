package main

import "fmt"

func main() {
	fmt.Println("Let's do if else!")

	loginCount := 23
	var result string
	if loginCount < 10 {	
		result = "Regular User"
	} else if loginCount < 20 {
		result = "Intermediate User"
	} else {
		result = "Super User"
	}	
	fmt.Println(result)

	if 9%2 == 0{
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num := 9; num < 10 {
		fmt.Println("Num is less than 10")
	} else if num > 10 {
		fmt.Println("Num is greater than 10")
	} else {
		fmt.Println("Num is equal to 10")
	}
}
