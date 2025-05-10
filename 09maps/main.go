package main

import "fmt"

func main() {
	fmt.Println("Let's do maps!")

	languages := make(map[string]string)

	languages["ruby"] = "Ruby"
	languages["rust"] = "Rust"
	languages["cpp"] = "Cpp"
	languages["rust"] = "C"

	// delete(languages, "rust")

	for key, value := range languages {
		fmt.Printf("for key %v, the value is %v\n", key, value)
	}

	fmt.Print(languages)

	// In maps, the key is unique. If you try to add a duplicate key, it will overwrite the previous value.
	// For example, if you add "rust" again, it will overwrite the previous value.
	// So, the final value of "rust" will be "C" instead of "Rust".
	// This is different from slices, where you can have duplicate values.
	// Maps are unordered, so the order of the keys may not be the same as the order in which you added them.
	// Maps are also not thread-safe, so you need to use mutexes or channels to synchronize access to them if you're using them in a concurrent program.
}
