package main

import (
	"fmt"
)

func main() {
	// var v string=nil ==> not allowed, Nil is not a type but a reserved word, you cannot use it in assignment.
	var data *string = nil
	if data == nil {
		fmt.Println(data)
	}
}
