package main

import "fmt"

func main(){
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["py"] = "Python"
	fmt.Printf("Type: %T\n",languages)
	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages,"RB")
	fmt.Println(languages)
	fmt.Println("\nLanguages")

	for key,value := range languages{
		fmt.Printf("Key: %v Value: %v\n",key,value)
	}
	fmt.Println("\nLanguages")

	for key := range languages{
		fmt.Printf("Key: %v\n",key)
	}
}