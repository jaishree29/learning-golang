package main

import "fmt"

func main() {
	fmt.Println("Let's do structs!")

	// No inhertance in Go, no super or parent
	Jaishree := User{"Jaishree", "jaishree@gmail.com", true, 22}

	fmt.Println(Jaishree.Name)
	fmt.Printf("Details: %+v\n", Jaishree)

	// If need all the details with parameters use %+v
	// If need only the values use %v

	
}

type User struct {
	Name	 string
	Email	 string
	Status	 bool
	Age 	 int
}