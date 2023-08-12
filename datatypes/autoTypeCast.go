package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 13 / 3.5
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))
}
