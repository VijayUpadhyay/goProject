package main

import "fmt"

type myInterface interface{}

type myStruct struct{}

func (ms myStruct) String() string {
	return "myStruct"
}

func main() {
	var i myInterface = myStruct{}
	fmt.Println(i)
}
