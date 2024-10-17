package main

import "fmt"

func namer(age int, names ...string) (string, int) {
	var name string
	for _, i := range names {
		name += i + " "
	}
	return name, age
}

func main() {
	fmt.Println("hello: ", func(name string) string {
		return name
	}("Aaryan"))

	names := []string{"Aaryan", "Thakur"}
	var ages int = 12
	name, age := namer(ages, names...)
	fmt.Println(name, age)
}
