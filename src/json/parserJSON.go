package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Index  int      `json:"index"`
	Fruits []string `json:"fruits"`
}

type response2 struct {
	Name   string   `json:"name"`
	Fruits []string `json:"vegetables"`
}

func main() {

	//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//byt1 := []byte(`{"index": 6, "fruits": ["apple","peach"]}`)
	byt2 := []byte(`{"name": "HOME", "vegetables": ["tomato","potato"]}`)

	byt := byt2

	var r *response2
	unmarshalResponse2(byt, r)

	unmarshalGeneric(byt)
}

func unmarshalGeneric(byt []byte) {
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("****unmarshalGeneric******")
	fmt.Printf("dat%+v", dat)

	//fmt.Println("false")
	// v := reflect.ValueOf(dat)
	//t := v.Type()
	// fmt.Printf("\nType%+v\n", v)

}
func unmarshalResponse2(byt []byte, r2 *response2) {
	if err := json.Unmarshal(byt, &r2); err != nil {
		panic(err)
	}
	fmt.Println("*****unmarshalResponse2*****")
	fmt.Printf("dat%+v\n", r2)

}
