package main

import "fmt"

type Person struct {
	Name      string
	Age       int64
	ExtraInfo map[string]interface{}
}

func main() {
	// map
	// make a map with capacity of 3
	m := make(map[string]int, 3)

	// add elements to the map
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println("Map:", m) // Output: map[one:1 two:2 three:3]

	// Len returns the number of elements in the map
	fmt.Println("Length of map:", len(m)) // Output: 3

	m1 := Person{
		Name:      "ZhangSan",
		Age:       18,
		ExtraInfo: make(map[string]interface{}),
	}

	m1.ExtraInfo["address"] = "Shanghai"

	fmt.Println(m1)
	fmt.Println(m1.ExtraInfo["address"])

	// Iterate through the map using range
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Iterate through the ExtraInfo map of m1 using range
	for key, value := range m1.ExtraInfo {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}

	delete(m, "one")
	fmt.Println("Map after deletion:", m) // Output: map[two:2 three:3]
}
