package Basics

import "fmt"
func main() {
	fmt.Println("vijay upadhyay")

	a:=make([]int,1)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println("slice")

	s,t:=slice(5,9)
	fmt.Println(s)
	fmt.Println(t)
	//fmt.Println(slice(5,8))
}

func slice(x,y int) (p,q int)  {
	s:= x-y
	t:= x+y
	return s,t
}