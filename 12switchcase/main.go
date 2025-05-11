package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	fmt.Println("Starting switch cases")

	// rand.NewSource(time.Now().UnixNano()) 
	
	// rand gives random numbers
	// If we don't use seed, it will generate same random number each time.
	// time.Now().UnixNano() gives current time in nanoseconds which is always different.


	diceNumber := rand.Intn(6) + 1  //Range is non inclusive so we add 1 i.e. Intn(6) gives 0-5
	fmt.Println("Value of dice is,", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a 1")
	case 2:
		fmt.Println("You rolled a 2")
	case 3:
		fmt.Println("You rolled a 3")
	case 4:
		fmt.Println("You rolled a 4")
	case 5:
		fmt.Println("You rolled a 5")
	case 6:
		fmt.Println("You rolled a 6")
		fallthrough // fallthrough is used to execute the next case as well
	default:
		fmt.Println("You got another chance")
	}
}