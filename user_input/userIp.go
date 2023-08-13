package main

import "fmt"

func main() {
	var i int
	fmt.Println("Enter integer value")
	fmt.Scan(&i)
	fmt.Println("Input value: ", i)

	// https://www.programiz.com/golang/take-input
	var j, k int
	fmt.Print("Type two numbers: ")
	fmt.Scanf("%v %v", &j, &k)
	fmt.Println("Your numbers are:", i, "and", j)
}
