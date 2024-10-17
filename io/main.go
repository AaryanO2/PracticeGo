package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter number of pizzas")
	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	fmt.Println("your Order: " + num)
}
