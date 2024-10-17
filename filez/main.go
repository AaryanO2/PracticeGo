package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(filename string) {
	txt, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text: \n", string(txt))
}

func main() {
	content := "hello world from file"
	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length: ", length)
	file.Close()
	readFile("./myfile.txt")
}
