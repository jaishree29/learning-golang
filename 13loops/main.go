package main

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	// while loop
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// do while loop
	i = 0
	for {
		if i >= 10 {
			break
		}
		println(i)
		i++
	}
}