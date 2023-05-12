package main

import (
	"encoding/json"
	"fmt"
)

type StructData struct {
	name       string
	rollNo     int
	address    string
	fatherName string
	motherName string
}

func main() {
	var test = StructData{"Vijay", 4, "address unknown", "umesh", "indu"}
	body, err := json.Marshal(test)
	if err != nil {
		fmt.Println("Unable to marshall to JSON")
		return
	} else {
		fmt.Println(string(body))
	}

	//Unmarshall response
	res := StructData{}
	json.Unmarshal([]byte(string(body)), &res)
	fmt.Println(res)
}
