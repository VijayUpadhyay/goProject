package main

import "fmt"

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := incrementor() // next is a function returned by incrementor
	fmt.Println(next())   // prints 1
	fmt.Println(next())   // prints 2
	fmt.Println(next())   // prints 3
}
