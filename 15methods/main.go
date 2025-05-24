package main

import "fmt"

func main() {
	fmt.Println("Let's do structs!")

	// No inhertance in Go, no super or parent
	Jaishree := User{"Jaishree", "jaishree@gmail.com", true, 22}

	fmt.Println(Jaishree.Name)
	fmt.Printf("Details: %+v\n", Jaishree)

	Jaishree.GetStatus()

	Jaishree.NewEmail()
	fmt.Println("New email is", Jaishree.Email) // This will not change the original email
	// This is because we are passing the value of the struct to the method

	// If need all the details with parameters use %+v
	// If need only the values use %v

	
}

type User struct {
	Name	 string
	Email	 string
	Status	 bool
	Age 	 int
}

func (u User) GetStatus() {
	fmt.Println("is user active?", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("New email is", u.Email)
}