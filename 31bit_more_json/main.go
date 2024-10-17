package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type course struct {
	Name     string `json:"Coursename"`
	Price    int
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func Encodejson(courses []course) ([]byte, error) {
	finalJson, err := json.MarshalIndent(courses, "", " ")
	return finalJson, err
}

func Decodejson(data []byte) (course, error) {
	var temp course
	chkValid := json.Valid(data)
	if chkValid {
		fmt.Println("JSON VALID")
		json.Unmarshal(data, &temp)
		return temp, nil
	} else {
		return temp, errors.New("invalid json")
	}
}

func main() {
	courses := []course{
		{"ML", 1223, "123123", []string{"cool", "NEW"}},
		{"AI", 12, "113", []string{"NEW"}},
		{"JS", 12, "113", nil},
	}
	finalJson, err := Encodejson(courses)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

	tmpjson := []byte(`
	{
		"Coursename": "ML",
		"Price": 1223,
		"tags": [
		"cool",
		"NEW"
		]
	}
	`)
	decoded, err := Decodejson(tmpjson)
	if err != nil {
		fmt.Printf("ERROR FOUND: %v\n", err)
	} else {
		fmt.Printf("%#v\n", decoded)
	}

	// just want key value pair??

	var mydata map[string]interface{}
	json.Unmarshal(tmpjson, &mydata)
	fmt.Println("MAP: \n", mydata)

	for k, v := range mydata {
		fmt.Printf("KEY: %v Value: %v KEY-TYPE: %T VALUE-TYPE: %T\n", k, v, k, v)
	}
}
