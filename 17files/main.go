package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	content := "This will go in a new file"

	file, err := os.Create("./mylcogofile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	
	fmt.Printf("Wrote %d bytes to file\n", length)
	defer file.Close()

	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Printf("Text data in file %s:\n%s\n", filename, databyte)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}