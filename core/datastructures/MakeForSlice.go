package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[0] = 5
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println("------------------------")
	b := make([]int, 2, 5) // make(type,len,capability)
	b[0] = 10
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println("------------------------")
	c := a[1:3]
	fmt.Println(c)
	fmt.Println(len(c))
}
