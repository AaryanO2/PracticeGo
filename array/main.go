package main

import "fmt"


func main(){
	fmt.Println("Welcome")
	var fruits  = [5]string {"Kiwi"}
	fruits[0] = "apple"
	fruits[3] = "mango"
	fruits[4] = "Pear"
	fmt.Println("Fruits are: ",fruits," Length: ",len(fruits))

}