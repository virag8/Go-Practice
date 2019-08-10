package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func main() {
	fmt.Println("Hello World" + string(89))
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input1 := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	input2 := map[string]interface{}{
		"first_name": "Mitchell",
		"age":        91,
		"last_name":  "Johnson",
		"homecity":   "",
	}
	e := Employee{FirstName: "Mitch", Last_Name: "Johnson", Age: 45}
	str := convertstructTojson(e)
	convertjsontostruct(str)
	var result1 Person
	err := mapstructure.Decode(input1, &result1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result1)

	var result2 Employee
	err = mapstructure.Decode(input2, &result2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%#v\n", result2)
}

type Employee struct {
	FirstName string `json:"first_name"`
	Last_Name string `json:"last_name"`
	Age       int    `json:"age"`
	HomeCity  string `json:"home_city,omitempty"`
}

type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

func convertjsontostruct(json_string string) {
	var emp Employee
	err := json.Unmarshal([]byte(json_string), &emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n%#v\n", emp)
}

func convertstructTojson(emp Employee) string {
	e, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("firstName", string(e))
	return string(e)
}
