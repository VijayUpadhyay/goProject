package main

import "fmt"
func main() {
	var s= []int{2,3,4,56,6,7,89,9}
	fmt.Println(s)
	s=s[:0]
	fmt.Println(s)
	s= s[2:4]
	fmt.Println(s)
	s=s[0:]
	fmt.Println("without ",s)
	s=s[0:3]
	fmt.Println(s)

	s=s[0:5]
	fmt.Println(s)

	var arr []int
	fmt.Println(arr)
	if arr==nil {
			fmt.Println("Inside if for nil")
	}
}
