package main

import "fmt"

const LoginToken int = 20000 //Public 

func main() {
	var username string = "Jaishree"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.613283923848876
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default
	var randomVariable int
	fmt.Println(randomVariable)
	fmt.Printf("Variable is of type: %T \n", randomVariable)

	//implicit type
	var name = "Jaishree"
	fmt.Printf("Variable is of type: %T \n", name)

	//no var style
	numberOfUsers := 800 // The symbol := means we are declaring and initializing the variable at the same time
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)	
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}