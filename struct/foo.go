package main

import (
	"fmt"
)

type Foo struct {
	value int
}

func PassStruct(foo Foo) {
	foo.value = 1
}

func PassStructPointer(foo *Foo) {
	foo.value = 1
}

func main() {
	var foo Foo

	fmt.Printf("before PassStruct: %v\n", foo.value)
	PassStruct(foo)
	fmt.Printf("after PassStruct: %v\n", foo.value)

	fmt.Printf("before PassStructPointer: %v\n", foo.value)
	PassStructPointer(&foo)
	fmt.Printf("after PassStructPointer: %v\n", foo.value)
}
