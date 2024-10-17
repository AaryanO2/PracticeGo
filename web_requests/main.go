package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:8000"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("Response: \n%v\nType: \n%T\n", response, response)

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Printf("Content: \n\n%v", content)
}
