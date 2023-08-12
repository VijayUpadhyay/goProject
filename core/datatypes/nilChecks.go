package main

import "fmt"

func main() {
	var a interface{}
	v, ok := a.(int)
	fmt.Println(v, ok)
	//fmt.Println(nil.(int))
	fmt.Println(a.(int))
}
