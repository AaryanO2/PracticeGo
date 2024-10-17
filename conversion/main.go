package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')
	fmt.Printf("Rating: %s", rating)
	newrating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("New Rating: ", newrating+1)
	}

}
