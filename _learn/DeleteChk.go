package main

import "fmt"

func main() {
	var NumOfObjectToAddOrDelete = 3
	var intArray = []int{31,3,4,66,77}
	size := len(intArray)-NumOfObjectToAddOrDelete
	intArray = intArray[:size]
	fmt.Println(intArray)
}
