package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Foo struct {
	A int
	B string
}
type Bar struct {
	C int
	D string
}

func main() {
	var temp sync.Map
	foo := Foo{A: 5, B: "bar"}
	bar := Bar{C: 5, D: "bar"}
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", foo)), temp)
	if err != nil {
		return
	}
	fmt.Println(temp)
	err = json.Unmarshal([]byte(fmt.Sprintf("%v", bar)), temp)
	if err != nil {
		return
	}
	fmt.Println(temp)
}
