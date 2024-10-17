package main

import "fmt"

func printer() {
	 for i := 1; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defer fmt.Println(" World")
	fmt.Println("Hello")
	printer()
}
