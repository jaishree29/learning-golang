package main

import "fmt"

func main() {
	// // for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// // for loop with range
	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// for i := range days {
	// 	fmt.Println(i, days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and day is %v\n", index, day)
	// }

	// for loop with condition
	rougeValue := 1;
	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println(rougeValue)
		rougeValue++
	}

	lco:
	fmt.Print("This is a label")

	// // while loop
	// i := 0
	// for i < 10 {
	// 	println(i)
	// 	i++
	// }

	// // do while loop
	// i = 0
	// for {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	println(i)
	// 	i++
	// }
}