package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User1 struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Age    int     `json:"Age"`
	Social Social1 `json:"social"`
}
type Social1 struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	var jsonFile = "user2.json"
	byteValue, _ := os.ReadFile(jsonFile)

	// we initialize our Users array
	var users map[string]User1

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err := json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	fmt.Println(users["user2a"])
}
