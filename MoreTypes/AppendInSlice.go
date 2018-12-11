package main

import "fmt"

func main() {
	var arr []int
	s:= arr[0:]
	r:=[]int{7,8,9}
	s=append(s,0,1,2,3,4)
	s=append(s,r)
	fmt.Println("s is:",s)
	fmt.Println(arr)
}
