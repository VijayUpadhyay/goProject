package main

import (
	"encoding/json"
	"fmt"
	"structures"
)

func main() {
	var test = structures.StructData{"Vijay", 4, "address unknown", "umesh", "indu"}
	body, err := json.Marshal(test)
	if err != nil {
		fmt.Println("Unable to marshall to JSON")
		return
	} else {
		fmt.Println(string(body))
	}

	//Unmarshall response
	res := structures.StructData{}
	json.Unmarshal([]byte(string(body)), &res)
	fmt.Println(res)
}
