package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Length of fruit list is:", len(fruitList))

	var vegList = [4]string{"aalu", "gobhi", "matar", "tamatar"}
	fmt.Println("Vegie list is:", vegList)
	fmt.Println("Length of vegie list is:", len(vegList))
}