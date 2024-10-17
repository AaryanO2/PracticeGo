package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	currentDate := time.Date(2020, time.August, 21,10,12,22,1231,time.UTC)
	fmt.Println(currentDate)
	fmt.Println(currentDate.Format("01-02-2006 Monday"))
}
