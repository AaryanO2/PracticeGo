package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PerformPostJson() {
	const myurl = "http://localhost:8000/post"
	//fake payload
	requestBody := strings.NewReader(`
		{
			"Course":"Golang",
			"Platform":"Lco",
			"Price": 12

		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
	var cont strings.Builder
	cont.Write(content)
	fmt.Println(cont.String())
}

func main() {
	PerformPostJson()
}
