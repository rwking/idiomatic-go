package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	str := []byte(`Richard King`)
	fmt.Println(str)

	var p struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	err := json.Unmarshal([]byte(`{"name": "Rich King", "age": 37}`), &p)

	fmt.Println(p, err)
}