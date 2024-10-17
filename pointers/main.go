package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printer(name *string) {
	*name = strings.TrimSpace(*name)+ " "+ "Thakur" 
}

func main() {
	fmt.Println("Welcome to a class on pointers")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Name: ",name)
	printer(&name)
	fmt.Println("Name: ",name)
}
