package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	City string `json:"gorod"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Arlen", City: "Almaty", Age: 21}

	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	text := `{"name":"Dana","gorod":"Shymkent"}`

	var p2 Person
	json.Unmarshal([]byte(text), &p2)

	fmt.Println(p2.Name)
	fmt.Println(p2.City)
}
