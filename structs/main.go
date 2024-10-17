package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func (u User) Details() (string, string, int) {
	return u.name, u.email, u.age
}

func (u *User) Updatename(name string) {
	u.name = name
}

func main() {
	hitesh := User{"Hitesh", "HKT@gmail.com", 12}
	fmt.Println(hitesh)
	name, email, age := hitesh.Details()
	// goto endd
	fmt.Printf("Details:\n%v\n%v\n%v\n", name, age, email)
// endd:
	hitesh.Updatename("Hitesh Kumar")
	name, email, age = hitesh.Details()
	fmt.Printf("Details:\n%v\n%v\n%v\n", name, age, email)
}
