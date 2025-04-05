package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's do slices!")

	var fruitList = []string{"Apple", "Banana"}
	fmt.Printf("The datatype of fruit list is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Grapes", "Papaya")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) //Here we are slicing the array using append method
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 342
	highScores[2] = 126
	highScores[1] = 99
	highScores[3] = 654

	highScores = append(highScores, 123, 23, 322) //appends the values by adding more memory
	sort.Ints(highScores) //Sorts the integer in ascending order
	fmt.Println(highScores)

	// How to remove a value from slices based on index
	var courses = []string{"reactjs", "python", "javascript", "swift"}
	fmt.Println(courses)
	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	// the three dots at the ends are used when there are more arguments used than expected
	fmt.Println(courses)
}