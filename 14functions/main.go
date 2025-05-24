package main

import "fmt"

func main() {
	fmt.Println("Starting with functions...")
	greeter()
	result := adder(1,2)
	fmt.Println("Result of adder is:", result)
	resultofProAdder := proAdder(1,2,3,4,5)
	fmt.Println("Result of proAdder is:", resultofProAdder)
}

func greeter() {
	fmt.Println("Namaste")
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) int {
	total := 0
	for _,val := range values {
		total+= val
	}
	return total
}