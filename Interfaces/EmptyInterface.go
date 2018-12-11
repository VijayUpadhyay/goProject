package main

import "fmt"

func main()  {
	var i interface{}
	i=50
	printDetails(i)
	i="vijay"
	printDetails(i)
}

func printDetails(i interface{})  {
	fmt.Printf("value is: %v and type is: %T",i,i)
	fmt.Println()
}
