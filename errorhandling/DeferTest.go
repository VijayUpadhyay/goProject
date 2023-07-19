package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("main 1st")
	fmt.Println("main 2nd")
	fmt.Println("Going to call method without defer")
	deferTest()
}
func deferTest() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}
