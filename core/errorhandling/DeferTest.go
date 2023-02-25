package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("first")
	defer fmt.Println("third")
	defer fmt.Println("fourth")
	fmt.Println("second")
	fmt.Println("-----------Inside at()---------------")
	deferTest()
}
func deferTest() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}