package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declaration and assignment
	var a1 byte = 97
	var a2 byte = 'b'

	// Printing without character formatting
	fmt.Println(a1)
	fmt.Println(a2)

	// Printing with character formatting
	fmt.Printf("%c\n", a1)
	fmt.Printf("%c\n", a2)
	fmt.Printf("a's type: %s \n", reflect.TypeOf(a1))

	// Shorthand declaration and assignment
	b1 := 97
	b2 := 'b'
	b3 := '♬'

	// Printing the value, unicode equivalent and type of the variable
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", b1, b1, reflect.TypeOf(b1))
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", b2, b2, reflect.TypeOf(b2))
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", b3, b3, reflect.TypeOf(b3))
}
