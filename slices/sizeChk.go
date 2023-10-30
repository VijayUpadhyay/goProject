package main

import "fmt"

func main() {
	sliceSizeCheck()
}

func sliceSizeCheck() {
	var a [5]int = [5]int{0, 1, 2} // 0,1,2,0,0
	fmt.Println(len(a), cap(a), a)
	b := a[1:4]       //1,2,0
	b = append(b, 10) //b: 1,2,0,10, a: 0,1,2,0,0,10
	fmt.Println("a: ", len(a), cap(a), a)
	fmt.Println("b: ", len(b), cap(b), b)
	sliceSizeCheckViaMethod(b)
	fmt.Println("After function call")
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)
}

func sliceSizeCheckViaMethod(b []int) {
	b = append(b, 10)
	b = append(b, 10)
	b = append(b, 10)
}
