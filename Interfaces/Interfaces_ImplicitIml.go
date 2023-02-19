package main

import (
	"fmt"
)

type  InterfaceTest interface {
	Test() string
	M()
}

type StructTest struct {
	firstName,lastName string
}


func (v StructTest) Test() string  {
	return v.firstName+" "+v.lastName
}

func (t StructTest) M() {
	fmt.Println("Inside method M")
	fmt.Println(string(t.firstName))
}
func main() {
	var i InterfaceTest
	l:= StructTest{"vijay","upadhyay"}
	i=l
	fmt.Printf("%d",i.M)
	fmt.Println()
	fmt.Println(i.Test())
}
