package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6)+1
	fmt.Println(dice)
	switch(dice){
	case 1:
		fmt.Println(1,"-")
		fallthrough
	case 2:
		fmt.Println(2,"-")
	case 3:
		fmt.Println(3,"-")
	case 4:
		fmt.Println(4,"-")
	case 5:
		fmt.Println(5,"-")
	case 6:
		fmt.Println(6,"-")
		fallthrough
	default:
		fmt.Println("Huhhhh")		
	}
}