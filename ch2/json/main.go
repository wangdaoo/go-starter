package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	// create a Person struct
	person := Person{
		Name:    "John",
		Age:     30,
		Hobbies: []string{"reading", "running", "swimming"},
	}

	fmt.Printf("person: %v\n", person)

	// encode the Person struct to JSON
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// print the encoded JSON
	fmt.Println(string(jsonBytes))

	// decode the JSON to a Person struct
	var decodedPerson Person
	err = json.Unmarshal(jsonBytes, &decodedPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// print the decoded Person struct
	fmt.Printf("%+v\n", decodedPerson)
}
