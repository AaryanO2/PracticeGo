package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "http://localhost:8000/postform"

func main() {
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	//formdata
	data := url.Values{}
	data.Add("Course", "Golang")
	data.Add("Platform", "Lco")
	// data.Add("Price", 12)
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(content)
	var context strings.Builder
	context.Write(content)
	fmt.Println(context.String())
}
