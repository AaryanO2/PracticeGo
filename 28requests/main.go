package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformRequest()
}

func PerformRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)
	content, _ := io.ReadAll(response.Body)
	var resp strings.Builder
	bytecount, _ := resp.Write(content)
	fmt.Println("bytecount: ", bytecount)
	fmt.Println("resp: ", resp.String())
	fmt.Println("Content: ", string(content))
}
