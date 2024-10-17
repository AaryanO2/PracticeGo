package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com:8080/homepage?name=Aryan"

func main() {
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())



	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("Data:\n", qparams["name"])


	partsofurl := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/homepage",
		RawPath: "user=Aaryan",
	}
	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)
}
