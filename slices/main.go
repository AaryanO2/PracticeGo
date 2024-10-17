package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"Tomato"}
	fmt.Printf("Type: %T\n", fruits)
	fruits = append(fruits, "Bhindi", "Onion", "Potato")
	fmt.Println("Fruits: ", fruits)
	fruits = append(fruits[0:1], append([]string{"Cabbage"},fruits[1:]...)...)
	fmt.Println("Fruits: ", fruits)
	fruits = append(fruits[0:1], fruits[2:]...)
	fmt.Println("Fruits: ", fruits)

	prices := make([]int, 4)
	prices[0] = 12
	prices[1] = 121
	prices[2] = 32
	prices[3] = 54
	prices = append(prices, 99)
	fmt.Println(prices)
	fmt.Println(sort.IntsAreSorted(prices))
	sort.Ints(prices)
	fmt.Println(prices)
	fmt.Println(sort.IntsAreSorted(prices))
}
