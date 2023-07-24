package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "vijay"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i], " type: ", reflect.TypeOf(s[i]))
	}
	fmt.Println("Iterate using range")
	for _, v := range s {
		fmt.Println(v, " type: ", reflect.TypeOf(v))
	}
}
