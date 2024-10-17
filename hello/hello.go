package main

import (
	"bufio"
	"fmt"
	"os"
)

const Marks = 12

func main() {
	var name = "hi"
	name = "hello"
	fmt.Printf("Hello %s", name)

	var chk bool = true
	if chk {
		fmt.Println("Bye")
	}
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Read " + input)
}
