package Basics

import (
	"fmt"
	"reflect"
)
func main() {
	fmt.Println("Hello World!!")

	var p =10
	fmt.Sprintf("%T",p)
	fmt.Sprintf("%v",p)
	fmt.Println("type is: ",reflect.TypeOf(p))
	fmt.Println("value is: ",p)
	fmt.Println("Using printf method for typeOf")
	fmt.Printf("%T",p)
}
