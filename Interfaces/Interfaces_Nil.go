package main

import (
	"fmt"
)

type TestNilInterface interface {
	checkNil()
}

type Vertex_2 struct {
	name string
}

func (v Vertex_2) checkNil()  {
	fmt.Println("Inside check nil method")
	fmt.Println(v.name)
}

func main() {
	var TestInt_2 TestNilInterface =Vertex_2{"ram"}
	fmt.Println("1st call")
	fmt.Println(TestInt_2.checkNil)
	fmt.Println("2nd call")
	TestInt_2.checkNil()
}
