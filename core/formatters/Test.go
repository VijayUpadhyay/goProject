package Basics

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World!!")

	var p = 10
	s := fmt.Sprintf("Data type of p is: %T", p)
	fmt.Println(s)
	str := fmt.Sprintf("Value of p is: %v", p)
	fmt.Println(str)
	fmt.Println("type is: ", reflect.TypeOf(p))
	fmt.Println("value is: ", p)
	fmt.Println("Using printf method for typeOf")
	fmt.Printf("%T", p)
}
